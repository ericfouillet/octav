package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const ConferenceStaffStdSelectColumns = "conference_staff.oid, conference_staff.conference_id, conference_staff.user_id, conference_staff.sort_order, conference_staff.created_on, conference_staff.modified_on"
const ConferenceStaffTable = "conference_staff"

type ConferenceStaffList []ConferenceStaff

func (c *ConferenceStaff) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.ConferenceID, &c.UserID, &c.SortOrder, &c.CreatedOn, &c.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceStaffTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlConferenceStaffDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceStaffTable)
		stmt.WriteString(` SET conference_id = ?, user_id = ?, sort_order = ? WHERE oid = ?`)
		library.Register("sqlConferenceStaffUpdateByOIDKey", stmt.String())
	})
}

func (c *ConferenceStaff) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ConferenceStaff.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", c)
	}
	c.CreatedOn = time.Now()
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
	stmt.WriteString(ConferenceStaffTable)
	stmt.WriteString(` (conference_id, user_id, sort_order, created_on, modified_on) VALUES (?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.ConferenceID, c.UserID, c.SortOrder, c.CreatedOn, c.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.OID = lii
	return nil
}

func (c ConferenceStaff) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceStaff.Update`).BindError(&err)
		defer g.End()
	}
	if c.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, c.OID)
		}
		stmt, err := library.GetStmt("sqlConferenceStaffUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.ConferenceID, c.UserID, c.SortOrder, c.OID)
		return err
	}
	return errors.New("OID must be filled")
}

func (c ConferenceStaff) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt("sqlConferenceStaffDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	return errors.New("column OID must be filled")
}

func (v *ConferenceStaffList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceStaff
	if capacity > 0 {
		res = make([]ConferenceStaff, 0, capacity)
	} else {
		res = []ConferenceStaff{}
	}

	for rows.Next() {
		vdb := ConferenceStaff{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}
