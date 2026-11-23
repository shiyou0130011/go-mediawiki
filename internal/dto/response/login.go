package response

type LoginAPIResp struct {
	Login struct {
		Result string
		Token  string
	}
}