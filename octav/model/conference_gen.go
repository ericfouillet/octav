package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = pdebug.Enabled
var _ = time.Time{}
var _ = sql.ErrNoRows

type rawConference struct {
	ID                        string               `json:"id"`
	Title                     string               `json:"title" l10n:"true"`
	Description               string               `json:"description,omitempty" l10n:"true"`
	CFPLeadText               string               `json:"cfp_lead_text,omitempty" l10n:"true"`
	CFPPreSubmitInstructions  string               `json:"cfp_pre_submit_instructions,omitempty" l10n:"true"`
	CFPPostSubmitInstructions string               `json:"cfp_post_submit_instructions,omitempty" l10n:"true"`
	ContactInformation        string               `json:"contact_information,omitempty" l10n:"true"`
	CoverURL                  string               `json:"cover_url"`
	RedirectURL               string               `json:"redirect_url"`
	SeriesID                  string               `json:"series_id,omitempty"`
	Series                    *ConferenceSeries    `json:"series,omitempty" decorate:"true"`
	SubTitle                  string               `json:"sub_title" l10n:"true"`
	Slug                      string               `json:"slug"`
	FullSlug                  string               `json:"full_slug,omitempty"`
	Status                    string               `json:"status"`
	BlogFeedbackAvailable     bool                 `json:"blog_feedback_available"`
	TimetableAvailable        bool                 `json:"timetable_available"`
	Timezone                  string               `json:"timezone"`
	Dates                     ConferenceDateList   `json:"dates,omitempty"`
	Administrators            UserList             `json:"administrators,omitempty" decorate:"true"`
	Venues                    VenueList            `json:"venues,omitempty" decorate:"true"`
	FeaturedSpeakers          FeaturedSpeakerList  `json:"featured_speakers,omitempty" decorate:"true"`
	Sponsors                  SponsorList          `json:"sponsors,omitempty" decorate:"true"`
	SessionTypes              SessionTypeList      `json:"session_types,omitempty" decorate:"true"`
	Tracks                    TrackList            `json:"tracks,omitempty" decorate:"true"`
	ExternalResources         ExternalResourceList `json:"external_resources,omitempty"`
}

func (v Conference) MarshalJSON() ([]byte, error) {
	var raw rawConference
	raw.ID = v.ID
	raw.Title = v.Title
	raw.Description = v.Description
	raw.CFPLeadText = v.CFPLeadText
	raw.CFPPreSubmitInstructions = v.CFPPreSubmitInstructions
	raw.CFPPostSubmitInstructions = v.CFPPostSubmitInstructions
	raw.ContactInformation = v.ContactInformation
	raw.CoverURL = v.CoverURL
	raw.RedirectURL = v.RedirectURL
	raw.SeriesID = v.SeriesID
	raw.Series = v.Series
	raw.SubTitle = v.SubTitle
	raw.Slug = v.Slug
	raw.FullSlug = v.FullSlug
	raw.Status = v.Status
	raw.BlogFeedbackAvailable = v.BlogFeedbackAvailable
	raw.TimetableAvailable = v.TimetableAvailable
	raw.Timezone = v.Timezone
	raw.Dates = v.Dates
	raw.Administrators = v.Administrators
	raw.Venues = v.Venues
	raw.FeaturedSpeakers = v.FeaturedSpeakers
	raw.Sponsors = v.Sponsors
	raw.SessionTypes = v.SessionTypes
	raw.Tracks = v.Tracks
	raw.ExternalResources = v.ExternalResources
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *Conference) Load(tx *sql.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.Conference.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.Conference{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *Conference) FromRow(vdb *db.Conference) error {
	v.ID = vdb.EID
	v.Title = vdb.Title
	if vdb.CoverURL.Valid {
		v.CoverURL = vdb.CoverURL.String
	}
	if vdb.RedirectURL.Valid {
		v.RedirectURL = vdb.RedirectURL.String
	}
	v.SeriesID = vdb.SeriesID
	if vdb.SubTitle.Valid {
		v.SubTitle = vdb.SubTitle.String
	}
	v.Slug = vdb.Slug
	v.Status = vdb.Status
	v.BlogFeedbackAvailable = vdb.BlogFeedbackAvailable
	v.TimetableAvailable = vdb.TimetableAvailable
	v.Timezone = vdb.Timezone
	return nil
}

func (v *Conference) ToRow(vdb *db.Conference) error {
	vdb.EID = v.ID
	vdb.Title = v.Title
	vdb.CoverURL.Valid = true
	vdb.CoverURL.String = v.CoverURL
	vdb.RedirectURL.Valid = true
	vdb.RedirectURL.String = v.RedirectURL
	vdb.SeriesID = v.SeriesID
	vdb.SubTitle.Valid = true
	vdb.SubTitle.String = v.SubTitle
	vdb.Slug = v.Slug
	vdb.Status = v.Status
	vdb.BlogFeedbackAvailable = v.BlogFeedbackAvailable
	vdb.TimetableAvailable = v.TimetableAvailable
	vdb.Timezone = v.Timezone
	return nil
}
