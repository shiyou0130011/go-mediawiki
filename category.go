package mediawiki

import (
	"encoding/json"
)

// Unmarshall response from page edits for query action
type outerQuery struct {
	Continue struct {
		Cmcontinue string
		Continue   string
	}
	Query struct {
		Categorymembers []struct {
			Title string
		}
	}
}

// SearchByCategory find all pages list with given category.
func (m *MWApi) SearchByCategory(category string) (result []string) {
	query := map[string]string{
		"action":  "query",
		"format":  "json",
		"prop":    "categoryinfo",
		"list":    "categorymembers",
		"cmtitle": "category:" + category,
	}

	var contin, cmcontinue string

	for {
		var subQuery = map[string]string{}
		if contin != "" && cmcontinue != "" {
			subQuery["continue"] = contin
			subQuery["cmcontinue"] = cmcontinue
		}

		data, err := m.API(query, subQuery)
		o := outerQuery{}
		err = json.Unmarshal(data, &o)
		if err != nil {
			break
		}
		for _, queryPage := range o.Query.Categorymembers {
			result = append(result, queryPage.Title)
		}
		if o.Continue.Cmcontinue == "" {
			break
		} else {
			cmcontinue = o.Continue.Cmcontinue
			contin = o.Continue.Continue
		}
	}

	return
}
