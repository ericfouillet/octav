// DO NOT EDIT. Automatically generated by hsup
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-urlenc"
)

const MaxResponseSize = (1 << 20) * 2

var _ = bytes.MinRead
var _ = json.Decoder{}
var transportJSONBufferPool = sync.Pool{
	New: allocTransportJSONBuffer,
}

func allocTransportJSONBuffer() interface{} {
	return &bytes.Buffer{}
}

func getTransportJSONBuffer() *bytes.Buffer {
	return transportJSONBufferPool.Get().(*bytes.Buffer)
}

func releaseTransportJSONBuffer(buf *bytes.Buffer) {
	buf.Reset()
	transportJSONBufferPool.Put(buf)
}

type Client struct {
	Client   *http.Client
	Endpoint string
}

func New(s string) *Client {
	return &Client{
		Client:   &http.Client{},
		Endpoint: s,
	}
}

func (c *Client) AddConferenceAdmin(in *model.AddConferenceAdminRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.AddConferenceAdmin").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/admin/add")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) AddConferenceDates(in *model.AddConferenceDatesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.AddConferenceDates").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/dates/add")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) CreateConference(in *model.CreateConferenceRequest) (ret *model.Conference, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Conference
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateRoom(in *model.CreateRoomRequest) (ret *model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Room
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateSession(in *model.CreateSessionRequest) (ret *model.Session, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Session
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateUser(in *model.CreateUserRequest) (ret *model.User, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.User
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateVenue(in *model.CreateVenueRequest) (ret *model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.CreateVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Venue
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) DeleteConference(in *model.DeleteConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteConferenceAdmin(in *model.DeleteConferenceAdminRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteConferenceAdmin").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/admin/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteConferenceDates(in *model.DeleteConferenceDatesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteConferenceDates").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/dates/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteRoom(in *model.DeleteRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteSession(in *model.DeleteSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteUser(in *model.DeleteUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteVenue(in *model.DeleteVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.DeleteVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) ListConference(in *model.ListConferenceRequest) (ret []model.Conference, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload []model.Conference
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListRoom(in *model.ListRoomRequest) (ret []model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload []model.Room
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListSessionByConference(in *model.ListSessionByConferenceRequest) (ret []model.Session, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListSessionByConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/schedule/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload []model.Session
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListUser(in *model.ListUserRequest) (ret []model.User, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload []model.User
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListVenue(in *model.ListVenueRequest) (ret []model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.ListVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload []model.Venue
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) LookupConference(in *model.LookupConferenceRequest) (ret *model.Conference, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Conference
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupRoom(in *model.LookupRoomRequest) (ret *model.Room, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Room
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupSession(in *model.LookupSessionRequest) (ret *model.Session, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Session
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupUser(in *model.LookupUserRequest) (ret *model.User, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.User
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupVenue(in *model.LookupVenueRequest) (ret *model.Venue, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.LookupVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	if pdebug.Enabled {
		pdebug.Printf("GET to %s", u.String())
	}
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	jsonbuf := getTransportJSONBuffer()
	defer releaseTransportJSONBuffer(jsonbuf)
	_, err = io.Copy(jsonbuf, io.LimitReader(res.Body, MaxResponseSize))
	defer res.Body.Close()
	if pdebug.Enabled {
		if err != nil {
			pdebug.Printf("failed to read respons buffer: %s", err)
		} else {
			pdebug.Printf("response buffer: %s", jsonbuf)
		}
	}
	if err != nil {
		return nil, err
	}

	var payload model.Venue
	err = json.Unmarshal(jsonbuf.Bytes(), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) UpdateConference(in *model.UpdateConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateConference").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/conference/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateRoom(in *model.UpdateRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateRoom").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/room/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateSession(in *model.UpdateSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateSession").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/session/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateUser(in *model.UpdateUserRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateUser").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/user/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) UpdateVenue(in *model.UpdateVenueRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("client.UpdateVenue").BindError(&err)
		defer g.End()
	}
	u, err := url.Parse(c.Endpoint + "/v1/venue/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	if pdebug.Enabled {
		pdebug.Printf("POST to %s", u.String())
		pdebug.Printf("%s", buf.String())
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}
