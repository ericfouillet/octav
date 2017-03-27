package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const LocalizedStringStdSelectColumns = "localized_strings.oid, localized_strings.parent_id, localized_strings.parent_type, localized_strings.name, localized_strings.language, localized_strings.localized"
const LocalizedStringTable = "localized_strings"

type LocalizedStringList []LocalizedString

func (l *LocalizedString) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&l.OID, &l.ParentID, &l.ParentType, &l.Name, &l.Language, &l.Localized)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(LocalizedStringTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlLocalizedStringDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(LocalizedStringTable)
		stmt.WriteString(` SET parent_id = ?, parent_type = ?, name = ?, language = ?, localized = ? WHERE oid = ?`)
		library.Register("sqlLocalizedStringUpdateByOIDKey", stmt.String())
	})
}

func (l *LocalizedString) Create(tx *sql.Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.LocalizedString.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", l)
	}
	doIgnore := false
	for _, opt := range opts {
		switch opt.(type) {
		case insertIgnoreOption:
			doIgnore = true
		}
	}

	stmt := bytes.Buffer{}
	stmt.WriteString("INSERT ")
	if doIgnore {
		stmt.WriteString("IGNORE ")
	}
	stmt.WriteString("INTO ")
	stmt.WriteString(LocalizedStringTable)
	stmt.WriteString(` (parent_id, parent_type, name, language, localized) VALUES (?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), l.ParentID, l.ParentType, l.Name, l.Language, l.Localized)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	l.OID = lii
	return nil
}

func (l LocalizedString) Update(tx *sql.Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`LocalizedString.Update`).BindError(&err)
		defer g.End()
	}
	if l.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, l.OID)
		}
		stmt, err := library.GetStmt("sqlLocalizedStringUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(l.ParentID, l.ParentType, l.Name, l.Language, l.Localized, l.OID)
		return err
	}
	return errors.New("OID must be filled")
}

func (l LocalizedString) Delete(tx *sql.Tx) error {
	if l.OID != 0 {
		stmt, err := library.GetStmt("sqlLocalizedStringDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(l.OID)
		return err
	}

	return errors.New("column OID must be filled")
}

func (v *LocalizedStringList) FromRows(rows *sql.Rows, capacity int) error {
	var res []LocalizedString
	if capacity > 0 {
		res = make([]LocalizedString, 0, capacity)
	} else {
		res = []LocalizedString{}
	}

	for rows.Next() {
		vdb := LocalizedString{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
