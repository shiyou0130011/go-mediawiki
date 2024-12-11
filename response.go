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

type starData struct {
	Data string `json:"*"`
}

type parseResponse struct {
	Parse struct {
		Title         string     `json:"title"`
		PageID        int        `json:"pageid"`
		Categories    []starData `json:"categories"`
		HTML          starData   `json:"text"`     // 內文的 HTML code
		WikiText      starData   `json:"wikitext"` // 內文的 wiki code
		Links         []starData `json:"links"`
		Images        []string   `json:"images"`
		Templates     []starData `json:"templates"` // 使用的模板
		ExternalLinks []string   `json:"externallinks"`
		Sections      []struct {
			TocLevel   int    `json:"toclevel"`
			Level      string `json:"level"`
			Line       string `json:"line"`
			Number     string `json:"number"`
			Index      string `json:"index"`
			Fromtitle  string `json:"fromtitle"`
			ByteOffset int    `json:"byteoffset"`
			Anchor     string `json:"anchor"`
		} `json:"sections"`
		DisplayTitle string `json:"displaytitle"`
	} `json:"parse"`
}

type mwError struct {
	Error struct {
		Code string
		Info string
	}
}
