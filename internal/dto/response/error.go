package response

type ErrorResp struct {
	ErrorBody struct {
		Code string
		Info string
	} `json:"error"`
}

func (e *ErrorResp) IsEmpty() bool {
	return e.ErrorBody.Code == "" && e.ErrorBody.Info == ""
}

func (e *ErrorResp) Error() string {
	if e.ErrorBody.Code == "" && e.ErrorBody.Info == "" {
		return "no error"
	}
	if e.ErrorBody.Info == "" {
		return e.ErrorBody.Code
	}
	if e.ErrorBody.Code == "" {
		return e.ErrorBody.Info
	}
	return e.ErrorBody.Code + ": " + e.ErrorBody.Info
}
