package response

type EditPageAPIResp struct {
	Edit struct {
		Result   string
		PageId   int
		Title    string
		OldRevId int
		NewRevId int
	}
}