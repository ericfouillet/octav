package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"errors"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards
func (v *Venue) Create(tx *db.Tx, vdb *db.Venue, payload model.CreateVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Venue.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := payload.L10N.CreateLocalizedStrings(tx, "Venue", vdb.EID); err != nil {
		return err
	}
	return nil
}

func (v *Venue) Update(tx *db.Tx, vdb *db.Venue, payload model.UpdateVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Venue.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == "" {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := v.populateRowForUpdate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Update(tx); err != nil {
		return err
	}

	return payload.L10N.Foreach(func(l, k, x string) error {
		if pdebug.Enabled {
			pdebug.Printf("Updating l10n string for '%s' (%s)", k, l)
		}
		ls := db.LocalizedString{
			ParentType: "Venue",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *Venue) ReplaceL10NStrings(tx *db.Tx, m *model.Venue, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Venue.ReplaceL10NStrings")
		defer g.End()
	}
	rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "Venue", m.ID, lang)
	if err != nil {
		return err
	}

	var l db.LocalizedString
	for rows.Next() {
		if err := l.Scan(rows); err != nil {
			return err
		}

		switch l.Name {
		case "name":
			if pdebug.Enabled {
				pdebug.Printf("Replacing for key 'name'")
			}
			m.Name = l.Localized
		case "address":
			if pdebug.Enabled {
				pdebug.Printf("Replacing for key 'address'")
			}
			m.Address = l.Localized
		}
	}
	return nil
}

func (v *Venue) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Venue.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.Venue{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "Venue"); err != nil {
		return err
	}
	return nil
}

func (v *Venue) LoadList(tx *db.Tx, vdbl *db.VenueList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
