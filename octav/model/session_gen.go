package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"encoding/json"
	"github.com/builderscon/octav/octav/tools"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}

type rawSession struct {
	ID                string       `json:"id"`
	ConferenceID      string       `json:"conference_id"`
	RoomID            string       `json:"room_id,omitempty"`
	SpeakerID         string       `json:"speaker_id"`
	SessionTypeID     string       `json:"session_type_id"`
	Title             string       `json:"title" l10n:"true"`
	Abstract          string       `json:"abstract" l10n:"true"`
	Memo              string       `json:"memo"`
	StartsOn          time.Time    `json:"starts_on"`
	Duration          int          `json:"duration"`
	MaterialLevel     string       `json:"material_level"`
	Tags              TagString    `json:"tags,omitempty" assign:"convert"`
	Category          string       `json:"category,omitempty"`
	SpokenLanguage    string       `json:"spoken_language,omitempty"`
	SlideLanguage     string       `json:"slide_language,omitempty"`
	SlideSubtitles    string       `json:"slide_subtitles,omitempty"`
	SlideURL          string       `json:"slide_url,omitempty"`
	VideoURL          string       `json:"video_url,omitempty"`
	PhotoPermission   string       `json:"photo_permission"`
	VideoPermission   string       `json:"video_permission"`
	HasInterpretation bool         `json:"has_interpretation"`
	Status            string       `json:"status"`
	Confirmed         bool         `json:"confirmed"`
	Conference        *Conference  `json:"conference,omitempy" decorate:"true"`
	Room              *Room        `json:"room,omitempty" decorate:"true"`
	Speaker           *User        `json:"speaker,omitempty" decorate:"true"`
	SessionType       *SessionType `json:"session_type,omitempty" decorate:"true"`
}

func (v Session) MarshalJSON() ([]byte, error) {
	var raw rawSession
	raw.ID = v.ID
	raw.ConferenceID = v.ConferenceID
	raw.RoomID = v.RoomID
	raw.SpeakerID = v.SpeakerID
	raw.SessionTypeID = v.SessionTypeID
	raw.Title = v.Title
	raw.Abstract = v.Abstract
	raw.Memo = v.Memo
	raw.StartsOn = v.StartsOn
	raw.Duration = v.Duration
	raw.MaterialLevel = v.MaterialLevel
	raw.Tags = v.Tags
	raw.Category = v.Category
	raw.SpokenLanguage = v.SpokenLanguage
	raw.SlideLanguage = v.SlideLanguage
	raw.SlideSubtitles = v.SlideSubtitles
	raw.SlideURL = v.SlideURL
	raw.VideoURL = v.VideoURL
	raw.PhotoPermission = v.PhotoPermission
	raw.VideoPermission = v.VideoPermission
	raw.HasInterpretation = v.HasInterpretation
	raw.Status = v.Status
	raw.Confirmed = v.Confirmed
	raw.Conference = v.Conference
	raw.Room = v.Room
	raw.Speaker = v.Speaker
	raw.SessionType = v.SessionType
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return tools.MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *Session) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.Session.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.Session{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(vdb); err != nil {
		return err
	}
	return nil
}

func (v *Session) FromRow(vdb db.Session) error {
	v.ID = vdb.EID
	v.ConferenceID = vdb.ConferenceID
	if vdb.RoomID.Valid {
		v.RoomID = vdb.RoomID.String
	}
	v.SpeakerID = vdb.SpeakerID
	v.SessionTypeID = vdb.SessionTypeID
	if vdb.Title.Valid {
		v.Title = vdb.Title.String
	}
	if vdb.Abstract.Valid {
		v.Abstract = vdb.Abstract.String
	}
	if vdb.Memo.Valid {
		v.Memo = vdb.Memo.String
	}
	if vdb.StartsOn.Valid {
		v.StartsOn = vdb.StartsOn.Time
	}
	v.Duration = vdb.Duration
	if vdb.MaterialLevel.Valid {
		v.MaterialLevel = vdb.MaterialLevel.String
	}
	if vdb.Tags.Valid {
		v.Tags = TagString(vdb.Tags.String)
	}
	if vdb.Category.Valid {
		v.Category = vdb.Category.String
	}
	if vdb.SpokenLanguage.Valid {
		v.SpokenLanguage = vdb.SpokenLanguage.String
	}
	if vdb.SlideLanguage.Valid {
		v.SlideLanguage = vdb.SlideLanguage.String
	}
	if vdb.SlideSubtitles.Valid {
		v.SlideSubtitles = vdb.SlideSubtitles.String
	}
	if vdb.SlideURL.Valid {
		v.SlideURL = vdb.SlideURL.String
	}
	if vdb.VideoURL.Valid {
		v.VideoURL = vdb.VideoURL.String
	}
	if vdb.PhotoPermission.Valid {
		v.PhotoPermission = vdb.PhotoPermission.String
	}
	if vdb.VideoPermission.Valid {
		v.VideoPermission = vdb.VideoPermission.String
	}
	v.HasInterpretation = vdb.HasInterpretation
	v.Status = vdb.Status
	v.Confirmed = vdb.Confirmed
	return nil
}

func (v *Session) ToRow(vdb *db.Session) error {
	vdb.EID = v.ID
	vdb.ConferenceID = v.ConferenceID
	vdb.RoomID.Valid = true
	vdb.RoomID.String = v.RoomID
	vdb.SpeakerID = v.SpeakerID
	vdb.SessionTypeID = v.SessionTypeID
	vdb.Title.Valid = true
	vdb.Title.String = v.Title
	vdb.Abstract.Valid = true
	vdb.Abstract.String = v.Abstract
	vdb.Memo.Valid = true
	vdb.Memo.String = v.Memo
	vdb.StartsOn.Valid = true
	vdb.StartsOn.Time = v.StartsOn
	vdb.Duration = v.Duration
	vdb.MaterialLevel.Valid = true
	vdb.MaterialLevel.String = v.MaterialLevel
	vdb.Tags.Valid = true
	vdb.Tags.String = string(v.Tags)
	vdb.Category.Valid = true
	vdb.Category.String = v.Category
	vdb.SpokenLanguage.Valid = true
	vdb.SpokenLanguage.String = v.SpokenLanguage
	vdb.SlideLanguage.Valid = true
	vdb.SlideLanguage.String = v.SlideLanguage
	vdb.SlideSubtitles.Valid = true
	vdb.SlideSubtitles.String = v.SlideSubtitles
	vdb.SlideURL.Valid = true
	vdb.SlideURL.String = v.SlideURL
	vdb.VideoURL.Valid = true
	vdb.VideoURL.String = v.VideoURL
	vdb.PhotoPermission.Valid = true
	vdb.PhotoPermission.String = v.PhotoPermission
	vdb.VideoPermission.Valid = true
	vdb.VideoPermission.String = v.VideoPermission
	vdb.HasInterpretation = v.HasInterpretation
	vdb.Status = v.Status
	vdb.Confirmed = v.Confirmed
	return nil
}
