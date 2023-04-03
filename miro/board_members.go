package miro

type BoardMembersService struct {
	client *Client
}

type ShareBoardRequest struct {
	Emails []string `json:"emails"`
}

//func (b *BoardMembersService) Share(id string, request *ShareBoardRequest) (*ListBoardsResponse, error) {
//	response := &ListBoardsResponse{}
//
//	url := fmt.Sprintf("%s/boards/%s/members", b.client.BaseURL, id)
//	err := b.client.Post(url, request, response)
//
//	return response, err
//}
