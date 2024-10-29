package mediawiki

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

type tokenResponse struct {
	Query struct {
		Tokens struct {
			Csrftoken string
		}
	}
}

type uploadResponse struct {
	Upload struct {
		Result string
	}
}

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

type parseResponse struct {
	Parse struct {
		Title      string `json:"title"`
		PageID     int    `json:"pageid"`
		Categories []struct {
			Name string `json:"*"`
		}
	}
}

type mwError struct {
	Error struct {
		Code string
		Info string
	}
}
