package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

var _ = time.Time{}

func (v *User) LookupFromPayload(tx *db.Tx, m *model.User, payload model.LookupUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.User from database")
	}
	if err := v.Decorate(tx, m, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.User from database")
	}
	return nil
}
func (v *User) Lookup(tx *db.Tx, m *model.User, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.User{}
	if err = r.Load(tx, id); err != nil {
		return errors.Wrap(err, "failed to load model.User from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *User) Create(tx *db.Tx, vdb *db.User, payload model.CreateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := payload.L10N.CreateLocalizedStrings(tx, "User", vdb.EID); err != nil {
		return err
	}
	return nil
}

func (v *User) Update(tx *db.Tx, vdb *db.User, payload model.UpdateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.Update (%s)", vdb.EID).BindError(&err)
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
			ParentType: "User",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *User) ReplaceL10NStrings(tx *db.Tx, m *model.User, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.User.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	if lang == "all" {
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ?`, "User", m.ID)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}
			if pdebug.Enabled {
				pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
			}
			m.L10N.Set(l.Language, l.Name, l.Localized)
		}
	} else {
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "User", m.ID, lang)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}

			switch l.Name {
			case "first_name":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'first_name'")
				}
				m.FirstName = l.Localized
			case "last_name":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'last_name'")
				}
				m.LastName = l.Localized
			}
		}
	}
	return nil
}

func (v *User) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("User.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.User{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "User"); err != nil {
		return err
	}
	return nil
}

func (v *User) LoadList(tx *db.Tx, vdbl *db.UserList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
