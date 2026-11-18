package mediawiki

import "github.com/shiyou0130011/go-mediawiki/internal/dto/response"

// Unmarshal login data...
type outerLogin struct {
	Login struct {
		Result string
		Token  string
	}
}

// Unmarshall response from page edits...
type outerEdit struct {
	Edit struct {
		Result   string
		PageId   int
		Title    string
		OldRevId int
		NewRevId int
	}
}

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

type starData struct {
	Data string `json:"*"`
}

type parseResponse = response.ParseAPIResp

type mwError struct {
	Error struct {
		Code string
		Info string
	}
}
