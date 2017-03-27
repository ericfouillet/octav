package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/builderscon/octav/octav/cache"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}
var _ = cache.WithExpires(time.Minute)
var _ = context.Background
var _ = errors.Wrap
var _ = model.Conference{}
var _ = db.Conference{}
var _ = sql.ErrNoRows
var _ = pdebug.Enabled

var conferenceSvc ConferenceSvc
var conferenceOnce sync.Once

func Conference() *ConferenceSvc {
	conferenceOnce.Do(conferenceSvc.Init)
	return &conferenceSvc
}

func (v *ConferenceSvc) LookupFromPayload(ctx context.Context, tx *sql.Tx, m *model.Conference, payload *model.LookupConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(ctx, tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.Conference from database")
	}
	if err := v.Decorate(ctx, tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.Conference from database")
	}
	return nil
}

func (v *ConferenceSvc) Lookup(ctx context.Context, tx *sql.Tx, m *model.Conference, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.Lookup").BindError(&err)
		defer g.End()
	}

	var r model.Conference
	c := Cache()
	key := c.Key("Conference", id)
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.Conference from database")
		}
		return &r, nil
	}, cache.WithExpires(time.Hour))
	if pdebug.Enabled {
		cacheSt := `HIT`
		if cacheMiss {
			cacheSt = `MISS`
		}
		pdebug.Printf(`CACHE %s: %s`, cacheSt, key)
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *ConferenceSvc) Create(ctx context.Context, tx *sql.Tx, vdb *db.Conference, payload *model.CreateConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx, payload.DatabaseOptions...); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	if err := payload.LocalizedFields.CreateLocalizedStrings(tx, "Conference", vdb.EID); err != nil {
		return errors.Wrap(err, `failed to populate localized strings`)
	}
	return nil
}

func (v *ConferenceSvc) Update(tx *sql.Tx, vdb *db.Conference) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	c := Cache()
	key := c.Key("Conference", vdb.EID)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	cerr := c.Delete(key)
	if pdebug.Enabled {
		if cerr != nil {
			pdebug.Printf(`CACHE ERR: %s`, cerr)
		}
	}
	return nil
}

func (v *ConferenceSvc) ReplaceL10NStrings(tx *sql.Tx, m *model.Conference, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	ls := LocalizedString()
	list := make([]db.LocalizedString, 0, 7)
	switch lang {
	case "", "en":
		if len(m.Title) > 0 && len(m.Description) > 0 && len(m.CFPLeadText) > 0 && len(m.CFPPreSubmitInstructions) > 0 && len(m.CFPPostSubmitInstructions) > 0 && len(m.ContactInformation) > 0 && len(m.SubTitle) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "Conference", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "title":
					if len(m.Title) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'title' (fallback en -> %s", l.Language)
						}
						m.Title = l.Localized
					}
				case "description":
					if len(m.Description) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'description' (fallback en -> %s", l.Language)
						}
						m.Description = l.Localized
					}
				case "cfp_lead_text":
					if len(m.CFPLeadText) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'cfp_lead_text' (fallback en -> %s", l.Language)
						}
						m.CFPLeadText = l.Localized
					}
				case "cfp_pre_submit_instructions":
					if len(m.CFPPreSubmitInstructions) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'cfp_pre_submit_instructions' (fallback en -> %s", l.Language)
						}
						m.CFPPreSubmitInstructions = l.Localized
					}
				case "cfp_post_submit_instructions":
					if len(m.CFPPostSubmitInstructions) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'cfp_post_submit_instructions' (fallback en -> %s", l.Language)
						}
						m.CFPPostSubmitInstructions = l.Localized
					}
				case "contact_information":
					if len(m.ContactInformation) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'contact_information' (fallback en -> %s", l.Language)
						}
						m.ContactInformation = l.Localized
					}
				case "sub_title":
					if len(m.SubTitle) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'sub_title' (fallback en -> %s", l.Language)
						}
						m.SubTitle = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "Conference", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				if pdebug.Enabled {
					pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
				}
				m.LocalizedFields.Set(l.Language, l.Name, l.Localized)
			}
		}
	default:
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "Conference", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "title":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'title'")
					}
					m.Title = l.Localized
				case "description":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'description'")
					}
					m.Description = l.Localized
				case "cfp_lead_text":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'cfp_lead_text'")
					}
					m.CFPLeadText = l.Localized
				case "cfp_pre_submit_instructions":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'cfp_pre_submit_instructions'")
					}
					m.CFPPreSubmitInstructions = l.Localized
				case "cfp_post_submit_instructions":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'cfp_post_submit_instructions'")
					}
					m.CFPPostSubmitInstructions = l.Localized
				case "contact_information":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'contact_information'")
					}
					m.ContactInformation = l.Localized
				case "sub_title":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'sub_title'")
					}
					m.SubTitle = l.Localized
				}
			}
		}
	}
	return nil
}

func (v *ConferenceSvc) Delete(tx *sql.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Conference.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.Conference{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return errors.Wrap(err, `failed to delete from database`)
	}
	c := Cache()
	key := c.Key("Conference", id)
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "Conference"); err != nil {
		return errors.Wrap(err, `failed to delete localized strings`)
	}
	return nil
}
