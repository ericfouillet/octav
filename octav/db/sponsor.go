package db

import (
	"database/sql"
	"strconv"

	"github.com/builderscon/octav/octav/tools"
)

func (v *SponsorList) LoadByConferenceSinceEID(tx *sql.Tx, confID, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Sponsor{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadByConferenceSince(tx, confID, s, limit)
}

func (v *SponsorList) LoadByConferenceSince(tx *sql.Tx, confID string, since int64, limit int) error {
	stmt := tools.GetBuffer()
	defer tools.ReleaseBuffer(stmt)

	stmt.WriteString(`SELECT `)
	stmt.WriteString(SponsorStdSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(SponsorTable)
	stmt.WriteString(` WHERE conference_id = ? AND `)
	stmt.WriteString(SponsorTable)
	stmt.WriteString(`.oid > ? ORDER BY oid ASC LIMIT `)
	stmt.WriteString(strconv.Itoa(limit))

	rows, err := tx.Query(stmt.String(), confID, since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}

func LoadSponsors(tx *sql.Tx, venues *SponsorList, cid string) error {
	stmt := tools.GetBuffer()
	defer tools.ReleaseBuffer(stmt)

	stmt.WriteString(`SELECT `)
	stmt.WriteString(SponsorStdSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(SponsorTable)
	stmt.WriteString(` WHERE `)
	stmt.WriteString(SponsorTable)
	stmt.WriteString(`.conference_id = ?`)
	stmt.WriteString(` ORDER BY sort_order ASC, group_name ASC`)

	rows, err := tx.Query(stmt.String(), cid)
	if err != nil {
		return err
	}

	var res SponsorList
	for rows.Next() {
		var u Sponsor
		if err := u.Scan(rows); err != nil {
			return err
		}

		res = append(res, u)
	}

	*venues = res
	return nil
}
