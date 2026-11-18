package response

type ContinueInfo struct {
		Cmcontinue string
		Continue   string
}

// Unmarshall response from page edits for query action
type CategorySearchResultResp struct {
	Continue ContinueInfo
	Query struct {
		Categorymembers []struct {
			Title string
		}
	}
}