package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"database/sql"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = pdebug.Enabled
var _ = time.Time{}
var _ = sql.ErrNoRows

func (v *Question) Load(tx *sql.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.Question.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.Question{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *Question) FromRow(vdb *db.Question) error {
	v.ID = vdb.EID
	v.SessionID = vdb.SessionID
	v.UserID = vdb.UserID
	v.Body = vdb.Body
	return nil
}

func (v *Question) ToRow(vdb *db.Question) error {
	vdb.EID = v.ID
	vdb.SessionID = v.SessionID
	vdb.UserID = v.UserID
	vdb.Body = v.Body
	return nil
}
