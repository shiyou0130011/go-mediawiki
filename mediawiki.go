package mediawiki

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"runtime"
	"strings"

	"github.com/shiyou0130011/go-mediawiki/internal/dto/response"
)

// MWApi is used to interact with the MediaWiki server.
type MWApi struct {
	username      string
	password      string
	Domain        string
	userAgent     string
	url           *url.URL
	client        *http.Client
	format        string
	edittoken     string
	UseBasicAuth  bool
	BasicAuthUser string
	BasicAuthPass string
}

// A Page represents a MediaWiki page and its metadata.
type Page = response.Page

// Helper function for translating MediaWiki errors in to Golang errors.
func checkError(resp []byte) error {
	var mwerror response.ErrorResp
	err := json.Unmarshal(resp, &mwerror)
	if err != nil {
		return nil
	}
	if mwerror.IsEmpty(){
		return nil
	} 
	return &mwerror
}

// New generates a new MediaWiki API (MWApi) struct.
func New(wikiURL string) (*MWApi, error) {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           cookiejar,
	}

	clientURL, err := url.Parse(wikiURL)
	if err != nil {
		return nil, err
	}

	return &MWApi{
		url:       clientURL,
		client:    &client,
		format:    "json",
		userAgent: "GO-Mediawiki/0.0.0 (Golang" + runtime.Version() + ") ",
	}, nil
}

// This will automatically add the user agent and encode the http request properly
func (m *MWApi) postForm(query url.Values) ([]byte, error) {
	request, err := http.NewRequest("POST", m.url.String(), strings.NewReader(query.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("user-agent", m.userAgent)
	if m.UseBasicAuth {
		request.SetBasicAuth(m.BasicAuthUser, m.BasicAuthPass)
	}
	resp, err := m.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = checkError(body); err != nil {
		return nil, err
	}

	return body, nil
}

// Login to the Mediawiki Website.
func (m *MWApi) Login(username, password string) error {
	if username == "" {
		return errors.New("empty username supplied")
	}
	if password == "" {
		return errors.New("empty password supplied")
	}
	m.username = username
	m.password = password

	query := map[string]string{
		"action":     "login",
		"lgname":     m.username,
		"lgpassword": m.password,
	}

	if m.Domain != "" {
		query["lgdomain"] = m.Domain
	}

	body, err := m.API(query)
	if err != nil {
		return err
	}

	var loginResp response.LoginAPIResp
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}

	if loginResp.Login.Result == "Success" {
		return nil
	} else if loginResp.Login.Result != "NeedToken" {
		return errors.New("Error logging in: " + loginResp.Login.Result)
	}

	// Need to use the login token
	query["lgtoken"] = loginResp.Login.Token

	body, err = m.API(query)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}

	if loginResp.Login.Result != "Success" {
		return errors.New("Error logging in: " + loginResp.Login.Result)
	}
	return nil
}

// GetEditToken retrieves an edit token from the MediaWiki site and saves it.
//
// This is necessary for editing any page.
//
// The Edit() function will call this automatically
// but it is available if you want to make direct
// calls to API().
func (m *MWApi) GetEditToken() error {
	query := map[string]string{
		"action":        "query",
		"meta":          "tokens",
		"formatversion": "2",
		"type":          "csrf",
	}

	body, err := m.API(query)
	if err != nil {
		return err
	}
	var tokenResponse response.TokenAPIResp
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return err
	}
	pl := tokenResponse.Query.Tokens.Csrftoken
	if pl == "" {
		return errors.New("no pages returned for edittoken query")
	}
	m.edittoken = pl
	return nil
}

// Logout of the MediaWiki website
func (m *MWApi) Logout() {
	m.API(map[string]string{"action": "logout"})
}

// Read returns the most recent revision of a Page. If an error occurs, nil is
// returned.
//
// Deprecated: This func was from sadbox/mediawiki. It is not used in this library. Please use “func (*MWApi) PageContent(string)” to get the content of a page.
func (m *MWApi) Read(pageName string) (*response.Page, error) {
	query := map[string]string{
		"action":  "query",
		"prop":    "revisions",
		"titles":  pageName,
		"rvlimit": "1",
		"rvprop":  "content|timestamp|user|comment",
	}
	body, err := m.API(query)
	if err != nil {
		return nil, err
	}

	var pageResponse response.PageAPIResponse
	err = json.Unmarshal(body, &pageResponse)
	if err != nil {
		return nil, err
	}

	if len(pageResponse.Query.Pages) != 1 {
		return nil, errors.New("received unexpected number of pages")
	}

	// we use a hacky way of extracting the map's lone value
	var page *response.Page
	for _, pg := range pageResponse.Query.Pages {
		page = &pg
	}
	return page, nil
}

// API is a generic interface to the Mediawiki API.
// Refer to the MediaWiki API reference for details.
//
// This is used by all internal functions to interact with the API.
func (m *MWApi) API(values ...map[string]string) ([]byte, error) {
	query := m.url.Query()
	for _, valuemap := range values {
		for key, value := range valuemap {
			query.Set(key, value)
		}
	}
	query.Set("format", m.format)
	body, err := m.postForm(query)
	if err != nil {
		return nil, err
	}
	return body, nil
}
