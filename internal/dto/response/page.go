package response

import "time"

type PageAPIResponse struct {
	Query struct {
		// The JSON response for this part of the struct is dumb.
		// It will return something like { '23': { 'pageid': 23 ...
		//
		// As a workaround you can use PageSlice which will create
		// a list of pages from the map.
		Pages map[string]Page
	}
}

// PageSlice generates a slice from Pages to work around the sillyness in
// the MediaWiki API.
func (r *PageAPIResponse) PageSlice() []Page {
	pl := []Page{}
	for _, page := range r.Query.Pages {
		pl = append(pl, page)
	}
	return pl
}

type Page struct {
	Pageid    int
	Ns        int
	Title     string
	Touched   string
	Lastrevid int
	// Mediawiki will return '' for zero, this makes me sad.
	// If for some reason you need this value you'll have to
	// do some type assertion sillyness.
	Counter   interface{}
	Length    int
	Edittoken string
	Revisions []struct {
		Revid         int       `json:"revid"`
		Parentid      int       `json:"parentid"`
		Minor         string    `json:"minor"`
		User          string    `json:"user"`
		Userid        int       `json:"userid"`
		Timestamp     time.Time `json:"timestamp"`
		Size          int       `json:"size"`
		Sha1          string    `json:"sha1"`
		ContentModel  string    `json:"contentmodel"`
		Comment       string    `json:"comment"`
		ParsedComment string    `json:"parsedcomment"`
		ContentFormat string    `json:"contentformat"`
		Body          string    `json:"*"` // Take note, MediaWiki literally returns { '*':
	}
	Imageinfo []struct {
		Url            string
		Descriptionurl string
	}
}
