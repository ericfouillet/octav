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

func (v *BlogEntry) Load(tx *sql.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.BlogEntry.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.BlogEntry{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *BlogEntry) FromRow(vdb *db.BlogEntry) error {
	v.ID = vdb.EID
	v.ConferenceID = vdb.ConferenceID
	v.Status = vdb.Status
	v.Title = vdb.Title
	v.URL = vdb.URL
	v.URLHash = vdb.URLHash
	return nil
}

func (v *BlogEntry) ToRow(vdb *db.BlogEntry) error {
	vdb.EID = v.ID
	vdb.ConferenceID = v.ConferenceID
	vdb.Status = v.Status
	vdb.Title = v.Title
	vdb.URL = v.URL
	vdb.URLHash = v.URLHash
	return nil
}
