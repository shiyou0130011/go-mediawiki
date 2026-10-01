package mediawiki

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type UploadConfig struct {
	FileComment string // Upload comment. Also used as the initial page text for new files if text is not specified.
	Text        string // Initial page text for the file.
	Summary     string // Upload summary.
}

// Upload a file
//
// This does a simple, but more error-prone upload. Mediawiki
// has a chunked upload version but it is only available in newer
// versions of the API.
//
// Automatically retrieves an edit token if necessary.
func (m *MWApi) Upload(dstFilename string, file io.Reader, config *UploadConfig) error {
	if m.edittoken == "" {
		err := m.GetEditToken()
		if err != nil {
			return err
		}
	}

	query := map[string]string{
		"action":   "upload",
		"filename": dstFilename,
		"token":    m.edittoken,
		"format":   m.format,
	}

	if config != nil {
		config = &UploadConfig{}
	}

	if config.FileComment != "" {
		query["comment"] = config.FileComment
	}
	if config.Text != "" {
		query["text"] = config.Text
	}
	if config.Summary != "" {
		query["text"] = config.Summary
	} else {
		query["summary"] = "Auto upload by go-mediawiki library."
	}

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	for key, value := range query {
		err := writer.WriteField(key, value)
		if err != nil {
			return err
		}
	}

	part, err := writer.CreateFormFile("file", dstFilename)
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", m.url.String(), buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("user-agent", m.userAgent)
	if m.UseBasicAuth {
		request.SetBasicAuth(m.BasicAuthUser, m.BasicAuthPass)
	}

	resp, err := m.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = checkError(body); err != nil {
		return err
	}

	var response uploadResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if !(response.Upload.Result == "Success" || response.Upload.Result == "Warning") {
		return errors.New(response.Upload.Result)
	}
	return nil
}
