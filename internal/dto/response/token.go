package response

type TokenAPIResp struct {
	Query struct {
		Tokens struct {
			Csrftoken string
		}
	}
}
