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

type rawFeaturedSpeaker struct {
	ID           string `json:"id"`
	ConferenceID string `json:"conference_id"`
	SpeakerID    string `json:"speaker_id"`
	AvatarURL    string `json:"avatar_url"`
	DisplayName  string `json:"display_name" l10n:"true"`
	Description  string `json:"description" l10n:"true"`
}

func (v FeaturedSpeaker) MarshalJSON() ([]byte, error) {
	var raw rawFeaturedSpeaker
	raw.ID = v.ID
	raw.ConferenceID = v.ConferenceID
	raw.SpeakerID = v.SpeakerID
	raw.AvatarURL = v.AvatarURL
	raw.DisplayName = v.DisplayName
	raw.Description = v.Description
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *FeaturedSpeaker) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.FeaturedSpeaker.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.FeaturedSpeaker{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *FeaturedSpeaker) FromRow(vdb *db.FeaturedSpeaker) error {
	v.ID = vdb.EID
	v.ConferenceID = vdb.ConferenceID
	if vdb.SpeakerID.Valid {
		v.SpeakerID = vdb.SpeakerID.String
	}
	if vdb.AvatarURL.Valid {
		v.AvatarURL = vdb.AvatarURL.String
	}
	v.DisplayName = vdb.DisplayName
	v.Description = vdb.Description
	return nil
}

func (v *FeaturedSpeaker) ToRow(vdb *db.FeaturedSpeaker) error {
	vdb.EID = v.ID
	vdb.ConferenceID = v.ConferenceID
	vdb.SpeakerID.Valid = true
	vdb.SpeakerID.String = v.SpeakerID
	vdb.AvatarURL.Valid = true
	vdb.AvatarURL.String = v.AvatarURL
	vdb.DisplayName = v.DisplayName
	vdb.Description = v.Description
	return nil
}
