package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"encoding/json"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = pdebug.Enabled
var _ = time.Time{}

type rawVenue struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name" l10n:"true" decorate:"true"`
	Address   string   `json:"address" l10n:"true" decorate:"true"`
	PlaceID   string   `json:"place_id,omitempty"`
	URL       string   `json:"url,omitempty"`
	Longitude float64  `json:"longitude,omitempty"`
	Latitude  float64  `json:"latitude,omitempty"`
	Rooms     RoomList `json:"rooms,omitempty"`
}

func (v Venue) MarshalJSON() ([]byte, error) {
	var raw rawVenue
	raw.ID = v.ID
	raw.Name = v.Name
	raw.Address = v.Address
	raw.PlaceID = v.PlaceID
	raw.URL = v.URL
	raw.Longitude = v.Longitude
	raw.Latitude = v.Latitude
	raw.Rooms = v.Rooms
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *Venue) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.Venue.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.Venue{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *Venue) FromRow(vdb *db.Venue) error {
	v.ID = vdb.EID
	v.Name = vdb.Name
	v.Address = vdb.Address
	v.PlaceID = vdb.PlaceID
	v.URL = vdb.URL
	v.Longitude = vdb.Longitude
	v.Latitude = vdb.Latitude
	return nil
}

func (v *Venue) ToRow(vdb *db.Venue) error {
	vdb.EID = v.ID
	vdb.Name = v.Name
	vdb.Address = v.Address
	vdb.PlaceID = v.PlaceID
	vdb.URL = v.URL
	vdb.Longitude = v.Longitude
	vdb.Latitude = v.Latitude
	return nil
}
