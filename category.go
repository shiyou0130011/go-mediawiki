package mediawiki

import (
	"encoding/json"
	"errors"
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

func (m *MWApi) PageCategoryList(title string) (result []string) {
	query := map[string]string{
		"action": "parse",
		"format": "json",
		"page":   title,
		"prop":   "categories",
	}
	b, err := m.API(query)
	if err != nil {
		return
	}

	var apiResult parseResponse

	if err := json.Unmarshal(b, &apiResult); err != nil {
		return
	}
	for _, c := range apiResult.Parse.Categories {
		result = append(result, c.Data)
	}
	return

}

func (m *MWApi) AddCategory(title, category string) error {
	pageCurrentCategoryList := m.PageCategoryList(title)
	for _, s := range pageCurrentCategoryList {
		if s == category {
			return errors.New("category is already exists")
		}
	}

	query := map[string]string{
		"title":      title,
		"action":     "edit",
		"format":     "json",
		"appendtext": "\n[[category:" + category + "]]",
	}
	_, err := m.API(query)
	return err
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
		if err != nil {
			break
		}
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
