package mediawiki

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Download a file.
//
// Returns a readcloser that must be closed manually. Refer to the
// example app for additional usage.
func (m *MWApi) Download(filename string) (io.ReadCloser, error) {
	// First get the direct url of the file
	query := map[string]string{
		"action": "query",
		"prop":   "imageinfo",
		"iiprop": "url",
		"titles": filename,
	}

	body, err := m.API(query)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	pl := response.PageSlice()

	if len(pl) < 1 {
		return nil, errors.New("no file found")
	}
	page := pl[0]
	if len(page.Imageinfo) < 1 {
		return nil, errors.New("no file found")
	}
	fileurl := page.Imageinfo[0].Url

	// Then return the body of the response
	request, err := http.NewRequest("GET", fileurl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("user-agent", m.userAgent)
	if m.UseBasicAuth {
		request.SetBasicAuth(m.BasicAuthUser, m.BasicAuthPass)
	}

	resp, err := m.client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
