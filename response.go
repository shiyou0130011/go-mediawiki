package mediawiki

import "github.com/shiyou0130011/go-mediawiki/internal/dto/response"

// Unmarshal login data...
type outerLogin = response.LoginAPIResp

// Unmarshall response from page edits...
type outerEdit = response.EditPageAPIResp

type tokenResponse = response.TokenAPIResp

type uploadResponse = response.UploadAPIResp

// Response is a struct used for unmarshaling the MediaWiki JSON response.
type Response struct {
	Query struct {
		// The JSON response for this part of the struct is dumb.
		// It will return something like { '23': { 'pageid': 23 ...
		//
		// As a workaround you can use PageSlice which will create
		// a list of pages from the map.
		Pages map[string]Page
	}
}

type starData = response.StarStringData

type parseResponse = response.ParseAPIResp

type mwError = response.ErrorResp
