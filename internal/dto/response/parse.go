package response


type ParseAPIResp struct {
	Parse struct {
		Title         string     `json:"title"`
		PageID        int        `json:"pageid"`
		Categories    []StarStringData `json:"categories"`
		HTML          StarStringData   `json:"text"`     // 內文的 HTML code
		WikiText      StarStringData   `json:"wikitext"` // 內文的 wiki code
		Links         []StarStringData `json:"links"`
		Images        []string   `json:"images"`
		Templates     []StarStringData `json:"templates"` // 使用的模板
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