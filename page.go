package mediawiki

import "encoding/json"

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

// SectionList will return all section title of page.
func (m *MWApi) SectionList(title string) (result []string, err error) {
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
