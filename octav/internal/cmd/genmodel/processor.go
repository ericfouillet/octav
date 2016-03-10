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
	pkgs, err := parser.ParseDir(fset, p.Dir, skipGenerated, parser.ParseComments)
	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		return errors.New("no packages to process...")
	}

	for _, pkg := range pkgs {
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

			if err := p.ProcessStruct(s); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Processor) ProcessStruct(s Struct) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessStruct %s", s.Name)
		defer g.End()
	}

	buf := bytes.Buffer{}

	varname := 'v'
	hasID := false

	buf.WriteString("// Automatically generated by genmodel utility. DO NOT EDIT!\n")
	buf.WriteString("package ")
	buf.WriteString(s.PackageName)
	buf.WriteString("\n\n")
	buf.WriteString("\nimport (\n")
	buf.WriteString("\n" + strconv.Quote("encoding/json"))
	buf.WriteString("\n" + strconv.Quote("github.com/builderscon/octav/octav/db"))
	buf.WriteString("\n" + strconv.Quote("github.com/lestrrat/go-pdebug"))
	buf.WriteString("\n)")

	fmt.Fprintf(&buf, "\n\nfunc (%c %s) GetPropNames() ([]string, error) {", varname, s.Name)
	fmt.Fprintf(&buf, "\nl, _ := %c.L10N.GetPropNames()", varname)
	buf.WriteString("\nreturn append(l, ")


	l10nfields := bytes.Buffer{}
	for _, f := range s.Fields {
		buf.WriteString(strconv.Quote(f.JSONName))
		buf.WriteString(",")
		if f.Name == "ID" {
			hasID = true
		}
		if f.L10N {
			l10nfields.WriteString(strconv.Quote(f.JSONName))
			l10nfields.WriteString(",")
		}
	}
	buf.WriteString("), nil")
	buf.WriteString("\n}")

	fmt.Fprintf(&buf, "\n\nfunc (%c %s) GetPropValue(s string) (interface{}, error) {", varname, s.Name)
	buf.WriteString("\nswitch s {")
	for _, f := range s.Fields {
		fmt.Fprintf(&buf, "\ncase %s:", strconv.Quote(f.JSONName))
		fmt.Fprintf(&buf, "\nreturn %c.%s, nil", varname, f.Name)
	}
	buf.WriteString("\ndefault:")
	fmt.Fprintf(&buf, "\nreturn %c.L10N.GetPropValue(s)", varname)
	buf.WriteString("\n}\n}")

	fmt.Fprintf(&buf, "\n\nfunc (%c %s) MarshalJSON() ([]byte, error) {", varname, s.Name)
	buf.WriteString("\nm := make(map[string]interface{})")
	for _, f := range s.Fields {
		fmt.Fprintf(&buf, "\nm[%s] = %c.%s", strconv.Quote(f.JSONName), varname, f.Name)
	}
	buf.WriteString("\nbuf, err := json.Marshal(m)")
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn nil, err")
	buf.WriteString("\n}")
	fmt.Fprintf(&buf, "\nreturn marshalJSONWithL10N(buf, %c.L10N)", varname)
	buf.WriteString("\n}")

	fmt.Fprintf(&buf, "\n\nfunc (%c *%s) UnmarshalJSON(data []byte) error {", varname, s.Name)
	buf.WriteString("\nm := make(map[string]interface{})")
	buf.WriteString("\nif err := json.Unmarshal(data, &m); err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")

	for _, f := range s.Fields {
		fmt.Fprintf(&buf, "\n\nif jv, ok := m[%s]; ok {", strconv.Quote(f.JSONName))
		buf.WriteString("\nswitch jv.(type) {")
		if strings.Contains(f.Type, "int") {
			buf.WriteString("\ncase float64:")
			fmt.Fprintf(&buf, "\n%c.%s = %s(jv.(float64))", varname, f.Name, f.Type)
		} else {
			fmt.Fprintf(&buf, "\ncase %s:", f.Type)
			fmt.Fprintf(&buf, "\n%c.%s = jv.(%s)", varname, f.Name, f.Type)
		}
		fmt.Fprintf(&buf, "\ndelete(m, %s)", strconv.Quote(f.JSONName))
		buf.WriteString("\ndefault:")
		fmt.Fprintf(&buf, "\nreturn ErrInvalidFieldType{Field: %s}", strconv.Quote(f.JSONName))
		buf.WriteString("\n}")
		buf.WriteString("\n}")
	}

	if l10nfields.Len() > 0 {
		fmt.Fprintf(&buf, "\n\nif err := ExtractL10NFields(m, &v.L10N, []string{%s}); err != nil {", l10nfields.String())
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
	}

	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	if hasID {
		fmt.Fprintf(&buf, "\n\nfunc (v *%s) Load(tx *db.Tx, id string) error {", s.Name)
		fmt.Fprintf(&buf, "\nvdb := db.%s{}", s.Name)
		buf.WriteString("\nif err := vdb.LoadByEID(tx, id); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}\n")
		buf.WriteString("\nif err := v.FromRow(vdb); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nif err := v.LoadLocalizedFields(tx); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")

		fmt.Fprintf(&buf, "\n\nfunc (v *%s) LoadLocalizedFields(tx *db.Tx) error {", s.Name)
		fmt.Fprintf(&buf, "\nls, err := db.LoadLocalizedStringsForParent(tx, v.ID, %s)", strconv.Quote(s.Name))
		buf.WriteString("\nif err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\n\nif len(ls) > 0 {")
		buf.WriteString("\nv.L10N = LocalizedFields{}")
		buf.WriteString("\nfor _, l := range ls {")
		buf.WriteString("\nv.L10N.Set(l.Language, l.Name, l.Localized)")
		buf.WriteString("\n}")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")

		fmt.Fprintf(&buf, "\n\nfunc (v *%s) FromRow(vdb db.%s) error {", s.Name, s.Name)
		buf.WriteString("\nv.ID = vdb.EID")
		for _, f := range s.Fields {
			if f.Name == "ID" {
				continue
			}
			fmt.Fprintf(&buf, "\nv.%s = vdb.%s", f.Name, f.Name)
		}
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")
	}

	fmt.Fprintf(&buf, "\n\nfunc (v *%s) Create(tx *db.Tx) error {", s.Name)
	if hasID {
		buf.WriteString("\n" + `if v.ID == "" {`)
		buf.WriteString("\nv.ID = UUID()")
		buf.WriteString("\n}")
	}
	fmt.Fprintf(&buf, "\n\nvdb := db.%s{", s.Name)
	for _, f := range s.Fields {
		if f.Name == "ID" {
			buf.WriteString("\nEID:     v.ID,")
		} else {
			fmt.Fprintf(&buf, "\n%s: v.%s,", f.Name, f.Name)
		}
	}
	buf.WriteString("\n}")
	buf.WriteString("\nif err := vdb.Create(tx); err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	if hasID {
		fmt.Fprintf(&buf, "\nif err := v.L10N.CreateLocalizedStrings(tx, %s, v.ID); err != nil {", strconv.Quote(s.Name))
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
	}
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	fmt.Fprintf(&buf, "\n\nfunc (v *%s) Delete(tx *db.Tx) error {", s.Name)
	buf.WriteString("\nif pdebug.Enabled {")
	fmt.Fprintf(&buf, "\n" + `g := pdebug.Marker("%s.Delete (%%s)", v.ID)`, s.Name)
	buf.WriteString("\ndefer g.End()")
	buf.WriteString("\n}")
	fmt.Fprintf(&buf, "\n\nvdb := db.%s{EID: v.ID}", s.Name)
	buf.WriteString("\nif err := vdb.Delete(tx); err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")
	fmt.Fprintf(&buf, "\nif err := db.DeleteLocalizedStringsForParent(tx, v.ID, %s); err != nil {", strconv.Quote(s.Name))
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	if hasID {
		fmt.Fprintf(&buf, "\n\nfunc (v *%sList) Load(tx *db.Tx, since string, limit int) error {", s.Name)
		fmt.Fprintf(&buf, "\nvdbl := db.%sList{}", s.Name)
		buf.WriteString("\nif err := vdbl.LoadSinceEID(tx, since, limit); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		fmt.Fprintf(&buf, "\nres := make([]%s, len(vdbl))", s.Name)
		buf.WriteString("\nfor i, vdb := range vdbl {")
		fmt.Fprintf(&buf, "\nv := %s{}", s.Name)
		buf.WriteString("\nif err := v.FromRow(vdb); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nif err := v.LoadLocalizedFields(tx); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nres[i] = v")
		buf.WriteString("\n}")
		buf.WriteString("\n*v = res")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")
	}

	fsrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return err
	}

	fn := filepath.Join(p.Dir, snakeCase(s.Name)+"_gen.go")
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
	cacheEnabled := true
	noScanner := false
	tablename := ""
	preCreate := ""
	postCreate := ""
	cacheExpires := "1800"

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

		if tablename == "" {
			tablename = fmt.Sprintf("%s_%s",
				ctx.Package,
				snakeCase(t.Name.Name),
			)
		}

		st := Struct{
			PackageName:  ctx.Package,
			CacheEnabled: cacheEnabled,
			CacheExpires: cacheExpires,
			Fields:       make([]StructField, 0, len(s.Fields.List)),
			Name:         t.Name.Name,
			NoScanner:    noScanner,
			PreCreate:    preCreate,
			PostCreate:   postCreate,
			Tablename:    tablename,
		}

	LoopFields:
		for _, f := range s.Fields.List {
			if len(f.Names) == 0 {
				continue
			}

			if unicode.IsLower(rune(f.Names[0].Name[0])) {
				continue
			}

			var jsname string
			var l10n bool
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

				st := reflect.StructTag(v)
				tag := st.Get("json")
				if tag == "-" {
					continue LoopFields
				}
				if tag == "" || tag[0] == ',' {
					jsname = f.Names[0].Name
				} else {
					tl := strings.SplitN(tag, ",", 2)
					jsname = tl[0]
				}

				tag = st.Get("l10n")
				if b, err := strconv.ParseBool(tag); err == nil && b {
					l10n = true
				}
			}

			typ, err := getTypeName(f.Type)
			if err != nil {
				return err
			}

			field := StructField{
				L10N: l10n,
				Name:     f.Names[0].Name,
				JSONName: jsname,
				Type:     typ,
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

