package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lestrrat/go-pdebug"
)

var ErrAnnotatedStructNotFound = errors.New("annotated struct was not found")

func snakeCase(s string) string {
	ret := []rune{}
	wasLower := false
	for len(s) > 0 {
		r, n := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			panic("yikes")
		}

		s = s[n:]
		if unicode.IsUpper(r) {
			if wasLower {
				ret = append(ret, '_')
			}
			wasLower = false
		} else {
			wasLower = true
		}

		ret = append(ret, unicode.ToLower(r))
	}
	return string(ret)
}

type Processor struct {
	Types []string
	Dir   string
}

func skipGenerated(fi os.FileInfo) bool {
	switch {
	case strings.HasSuffix(fi.Name(), "_gen.go"):
		return false
	case strings.HasSuffix(fi.Name(), "_gen.go"):
		return false
	}
	return true
}

func (p *Processor) Do() error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, filepath.Join(p.Dir, "model"), skipGenerated, parser.ParseComments)
	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		return errors.New("no packages to process...")
	}

	for _, pkg := range pkgs {
		if strings.HasSuffix(pkg.Name, "_test") {
			continue
		}

		if err := p.ProcessPkg(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) ShouldProceed(s Struct) bool {
	if len(p.Types) == 0 {
		return true
	}

	for _, t := range p.Types {
		if t == s.Name {
			return true
		}
	}
	return false
}

func (p *Processor) ProcessPkg(pkg *ast.Package) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessPkg %s", pkg.Name)
		defer g.End()
	}

	buf := bytes.Buffer{}
	buf.WriteString("package model")
	buf.WriteString("\n\n// Automatically generated by gentransport utility. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	buf.WriteString("\n" + strconv.Quote("encoding/json"))
	buf.WriteString("\n" + strconv.Quote("errors"))
	buf.WriteString("\n\n" + strconv.Quote("github.com/builderscon/octav/octav/tools"))
	buf.WriteString("\n\n" + strconv.Quote("github.com/lestrrat/go-urlenc"))
	buf.WriteString("\n)")

	for fn, f := range pkg.Files {
		pdebug.Printf("Checking file %s", fn)
		for _, s := range p.ExtractStructs(pkg, f) {
			if pdebug.Enabled {
				pdebug.Printf("Checking struct %s", s.Name)
			}
			if !p.ShouldProceed(s) {
				if pdebug.Enabled {
					pdebug.Printf("Skipping struct %s", s.Name)
				}
				continue
			}

			if err := p.ProcessStruct(&buf, s); err != nil {
				return err
			}
		}
	}

	fsrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return err
	}

	// TODO: we probably need to generate for multiple transport
	// messages, but at this point I'm a bit lazy to do all the
	// consolidation/
	fn := filepath.Join(p.Dir, "model", "transport_gen.go")
	if pdebug.Enabled {
		pdebug.Printf("Generating file %s", fn)
	}
	fi, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fi.Close()

	if _, err := fi.Write(fsrc); err != nil {
		return err
	}

	return nil
}

func (p *Processor) ProcessStruct(buf *bytes.Buffer, s Struct) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessStruct %s", s.Name)
		defer g.End()
	}

	data := map[string]string{
		"json":   "JSON",
		"urlenc": "URL",
	}

	fmt.Fprintf(buf, "\n\nfunc (r %s) collectMarshalData() map[string]interface{} {", s.Name)
	buf.WriteString("\nm := make(map[string]interface{})")
	for _, f := range s.Fields {
		if !f.Maybe {
			fmt.Fprintf(buf, "\nm[%s] = r.%s", strconv.Quote(f.JSONName), f.Name)
			/*
				if strings.HasPrefix(f.Type, "[]") {
					t := f.Type[2:]
					fmt.Fprintf(buf, "\nlist := make(%s, len(r.%s)", f.Type, f.Name)
					fmt.Fprintf(buf, "\nfor i, el := range r.%s {", f.Name)
					buf.WriteString("\nswitch el.(type) {")
					fmt.Fprintf(buf, "\ncase %s:", t)
					fmt.Fprintf(buf, "\nlist[i] = el.(%s)", t)
					buf.WriteString("\ndefault:")
					buf.
			*/
		} else {
			fmt.Fprintf(buf, "\nif r.%s.Valid() {", f.Name)
			fmt.Fprintf(buf, "\nm[%s] = r.%s.Value()", strconv.Quote(f.JSONName), f.Name)
			buf.WriteString("\n}")
		}
	}
	buf.WriteString("\nreturn m")
	buf.WriteString("\n}")

	for _, prefix := range []string{"json", "urlenc"} {
		method := data[prefix]
		fmt.Fprintf(buf, "\n\nfunc (r %s) Marshal%s() ([]byte, error) {", s.Name, method)
		buf.WriteString("\nm := r.collectMarshalData()")
		fmt.Fprintf(buf, "\nbuf, err := %s.Marshal(m)", prefix)
		buf.WriteString("\nif err != nil {")
		buf.WriteString("\nreturn nil, err")
		buf.WriteString("\n}")
		if !s.HasL10N {
			buf.WriteString("\nreturn buf, nil")
		} else {
			fmt.Fprintf(buf, "\nreturn tools.Marshal%sWithL10N(buf, r.L10N)", method)
		}
		buf.WriteString("\n}")
	}

	fmt.Fprintf(buf, "\n\nfunc (r *%s) UnmarshalJSON(data []byte) error {", s.Name)
	buf.WriteString("\nm := make(map[string]interface{})")
	buf.WriteString("\nif err := json.Unmarshal(data, &m); err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")
	buf.WriteString("\nreturn r.Populate(m)")
	buf.WriteString("\n}")

	fmt.Fprintf(buf, "\n\nfunc (r *%s) Populate(m map[string]interface{}) error {", s.Name)
	rx := regexp.MustCompile(`^(u?int\d*|float(32|64))`)
	for _, f := range s.Fields {
		fmt.Fprintf(buf, "\nif jv, ok := m[%s]; ok {", strconv.Quote(f.JSONName))
		if f.Maybe {
			fmt.Fprintf(buf, "\nif err := r.%s.Set(jv); err != nil {", f.Name)
			fmt.Fprintf(buf, "\n"+`return errors.New("set field %s failed: " + err.Error())`, f.Name)
			buf.WriteString("\n}")
			fmt.Fprintf(buf, "\ndelete(m, %s)", strconv.Quote(f.JSONName))
		} else if f.HasExtract {
			fmt.Fprintf(buf, "\nif err := r.%s.Extract(jv); err != nil {", f.Name)
			fmt.Fprintf(buf, "\n"+`return errors.New("extract field %s failed: " + err.Error())`, f.Name)
			buf.WriteString("\n}")
			fmt.Fprintf(buf, "\ndelete(m, %s)", strconv.Quote(f.JSONName))
		} else {
			buf.WriteString("\nswitch jv.(type) {")
			// XXX dirty hack. this should be fixed
			convert := ""
			typ := f.Type

			if !strings.HasPrefix(typ, "[]") {
				if rx.MatchString(f.Type) {
					convert = f.Type
					typ = "float64"
				}

				fmt.Fprintf(buf, "\ncase %s:", typ)
				if convert == "" {
					fmt.Fprintf(buf, "\nr.%s = jv.(%s)", f.Name, typ)
				} else {
					fmt.Fprintf(buf, "\nr.%s = %s(jv.(%s))", f.Name, convert, typ)
				}
				fmt.Fprintf(buf, "\ndelete(m, %s)", strconv.Quote(f.JSONName))
			} else {
				ltyp := typ[2:]
				buf.WriteString("\ncase []interface{}:")
				buf.WriteString("\njvl := jv.([]interface{})")
				fmt.Fprintf(buf, "\nlist := make(%s, len(jvl))", typ)
				buf.WriteString("\nfor i, el := range jvl {")
				buf.WriteString("\nswitch el.(type) {")
				fmt.Fprintf(buf, "\ncase %s:", ltyp)
				fmt.Fprintf(buf, "\nlist[i] = el.(%s)", ltyp)
				buf.WriteString("\ndefault:")
				fmt.Fprintf(buf, "\nreturn ErrInvalidJSONFieldType{Field:%s}", strconv.Quote(f.JSONName))
				buf.WriteString("\n}")
				buf.WriteString("\n}")
				fmt.Fprintf(buf, "\nr.%s = list", f.Name)
			}

			buf.WriteString("\ndefault:")
			fmt.Fprintf(buf, "\nreturn ErrInvalidJSONFieldType{Field:%s}", strconv.Quote(f.JSONName))
			buf.WriteString("\n}")
		}
		buf.WriteString("\n}")
	}
	if s.HasL10N {
		buf.WriteString("\nif err := tools.ExtractL10NFields(m, &r.L10N, []string{")
		for i, f := range s.Fields {
			buf.WriteString(strconv.Quote(f.JSONName))
			if i != len(s.Fields)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
	}
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	if s.HasL10N {
		fmt.Fprintf(buf, "\n\nfunc (r *%s) GetPropNames() ([]string, error) {", s.Name)
		buf.WriteString("\nl, _ := r.L10N.GetPropNames()")
		buf.WriteString("\nreturn append(l, ")
		for i, f := range s.Fields {
			buf.WriteString(strconv.Quote(f.JSONName))
			if i != len(s.Fields)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("), nil")
		buf.WriteString("\n}")

		fmt.Fprintf(buf, "\n\nfunc (r *%s) SetPropValue(s string, v interface{}) error {", s.Name)
		buf.WriteString("\nswitch s {")
		for _, f := range s.Fields {
			fmt.Fprintf(buf, "\ncase %s:", strconv.Quote(f.JSONName))
			if f.Maybe {
				fmt.Fprintf(buf, "\nreturn r.%s.Set(v)", f.Name)
			} else {
				fmt.Fprintf(buf, "\nif jv, ok := v.(%s); ok {", f.Type)
				fmt.Fprintf(buf, "\nr.%s = jv", f.Name)
				buf.WriteString("\nreturn nil")
				buf.WriteString("\n}")
			}
		}
		buf.WriteString("\ndefault:")
		buf.WriteString("\n" + `return errors.New("unknown column '" + s + "'")`)
		buf.WriteString("\n}")
		buf.WriteString("\nreturn ErrInvalidFieldType{Field: s}")
		buf.WriteString("\n}")
	}

	return nil
}

func (p *Processor) ExtractStructs(pkg *ast.Package, f *ast.File) []Struct {
	ctx := &InspectionCtx{
		Package: pkg.Name,
	}

	ast.Inspect(f, ctx.ExtractStructs)
	return ctx.Structs
}

func (c *InspectionCtx) ExtractStructs(n ast.Node) bool {
	var decl *ast.GenDecl
	var ok bool
	var err error

	if decl, ok = n.(*ast.GenDecl); !ok {
		return true
	}

	if decl.Tok != token.TYPE {
		return true
	}

	if err = c.ExtractStructsFromDecl(decl); err != nil {
		return true
	}

	return false
}

func (ctx *InspectionCtx) ExtractStructsFromDecl(decl *ast.GenDecl) error {
	for _, spec := range decl.Specs {
		var t *ast.TypeSpec
		var s *ast.StructType
		var ok bool

		if t, ok = spec.(*ast.TypeSpec); !ok {
			return ErrAnnotatedStructNotFound
		}

		if s, ok = t.Type.(*ast.StructType); !ok {
			return ErrAnnotatedStructNotFound
		}

		cgroup := decl.Doc
		if cgroup == nil {
			continue
		}
		istransport := false
		for _, c := range cgroup.List {
			if strings.HasPrefix(strings.TrimSpace(strings.TrimPrefix(c.Text, "//")), "+transport") {
				istransport = true
				break
			}
		}
		if !istransport {
			continue
		}

		st := Struct{
			PackageName: ctx.Package,
			Fields:      make([]StructField, 0, len(s.Fields.List)),
			Name:        t.Name.Name,
			HasL10N:     false,
		}

		for _, f := range s.Fields.List {
			if len(f.Names) == 0 {
				continue
			}

			if f.Names[0].Name == "L10N" {
				st.HasL10N = true
				continue
			}

			if unicode.IsLower(rune(f.Names[0].Name[0])) {
				continue
			}

			var jsname string
			var urlname string
			var l10n bool
			var hasExtract bool
			if f.Tag != nil {
				v := f.Tag.Value
				if len(v) >= 2 {
					if v[0] == '`' {
						v = v[1:]
					}
					if v[len(v)-1] == '`' {
						v = v[:len(v)-1]
					}
				}

				sf := reflect.StructTag(v)
				if tag := sf.Get("urlenc"); tag != "-" {
					if tag == "" || tag[0] == ',' {
						urlname = f.Names[0].Name
					} else {
						tl := strings.SplitN(tag, ",", 2)
						urlname = tl[0]
					}
				}

				if tag := sf.Get("json"); tag != "-" {
					if tag == "" || tag[0] == ',' {
						jsname = f.Names[0].Name
					} else {
						tl := strings.SplitN(tag, ",", 2)
						jsname = tl[0]
					}
				}

				if tag := sf.Get("l10n"); tag != "" {
					if b, err := strconv.ParseBool(tag); err == nil && b {
						st.HasL10N = true
						l10n = true
					}
				}

				if tag := sf.Get("extract"); tag == "true" {
					hasExtract = true
				}
			}

			if jsname == "-" || jsname == "" {
				continue
			}

			typ, err := getTypeName(f.Type)
			if err != nil {
				return err
			}

			var maybe bool
			if strings.HasPrefix(typ, "Maybe") {
				maybe = true
			}
			field := StructField{
				HasExtract: hasExtract,
				L10N:       l10n,
				Maybe:      maybe,
				Name:       f.Names[0].Name,
				JSONName:   jsname,
				URLName:    urlname,
				Type:       typ,
			}

			st.Fields = append(st.Fields, field)
		}
		ctx.Structs = append(ctx.Structs, st)
	}

	return nil
}

func getTypeName(ref ast.Expr) (string, error) {
	var typ string
	switch ref.(type) {
	case *ast.Ident:
		typ = ref.(*ast.Ident).Name
	case *ast.SelectorExpr:
		typ = ref.(*ast.SelectorExpr).Sel.Name
	case *ast.StarExpr:
		return getTypeName(ref.(*ast.StarExpr).X)
	case *ast.ArrayType:
		typ = "[]" + ref.(*ast.ArrayType).Elt.(*ast.Ident).Name
	case *ast.MapType:
		mt := ref.(*ast.MapType)
		typ = "map[" + mt.Key.(*ast.Ident).Name + "]" + mt.Value.(*ast.Ident).Name
	default:
		fmt.Printf("%#v\n", ref)
		return "", errors.New("field type not supported")
	}
	return typ, nil
}
