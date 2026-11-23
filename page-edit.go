package mediawiki

import (
	"encoding/json"
	"errors"

	"github.com/shiyou0130011/go-mediawiki/internal/dto/response"
)

// Edit a page.
//
// This function will request an edit token if the MWApi struct doesn't already
// contain one.
func (m *MWApi) Edit(values map[string]string) error {
	if m.edittoken == "" {
		err := m.GetEditToken()
		if err != nil {
			return err
		}
	}
	query := map[string]string{
		"action": "edit",
		"token":  m.edittoken,
	}
	body, err := m.API(query, values)
	if err != nil {
		return err
	}

	var response response.EditPageAPIResp
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Edit.Result != "Success" {
		return errors.New(response.Edit.Result)
	}
	return nil
}
