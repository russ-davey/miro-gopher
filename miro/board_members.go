package miro

type BoardMembersService struct {
	client *Client
}

type ShareBoardRequest struct {
	Emails []string `json:"emails"`
}
