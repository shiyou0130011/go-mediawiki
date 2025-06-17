package mediawiki

import (
	"encoding/json"
	"fmt"
)

// PageContent will return the content of a page in wikitext format.
func (m *MWApi) PageContent(title string) (string, error) {
	data, err := m.API(map[string]string{
		"action": "parse",
		"format": "json",
		"prop":   "wikitext",
		"page":   title,
	})
	if err != nil {
		return "", err
	}

	var result parseResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return "", err
	}

	return result.Parse.WikiText.Data, nil
}

// SectionContent will return the content of a specific section in a page.
func (m *MWApi) SectionContent(title string, sectionIndex int) (string, error) {
	data, err := m.API(map[string]string{
		"action":  "parse",
		"format":  "json",
		"prop":    "wikitext",
		"page":    title,
		"section": fmt.Sprint(sectionIndex),
	})
	if err != nil {
		return "", err
	}

	var result parseResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return "", err
	}

	return result.Parse.WikiText.Data, nil
}

// SectionTitleList will return all section title of page.
func (m *MWApi) SectionTitleList(title string) (result []string, err error) {
	data, err := m.API(map[string]string{
		"action": "parse",
		"format": "json",
		"prop":   "sections",
		"page":   title,
	})
	if err != nil {
		return
	}

	var respResult parseResponse
	if err = json.Unmarshal(data, &respResult); err != nil {
		return
	}

	// 因為每篇文章都會有沒標題的 top section，所以初始化時先建立 index 0 的值
	result = append(result, "")

	for _, data := range respResult.Parse.Sections {
		result = append(result, data.Line)
	}
	return
}
