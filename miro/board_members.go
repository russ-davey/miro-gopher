package miro

import (
	"fmt"
)

const (
	RoleViewer    = "viewer"
	RoleCommenter = "commenter"
	RoleEditor    = "editor"
	RoleCoOwner   = "coowner"
	RoleOwner     = "owner"
	RoleGuest     = "guest"
)

type BoardMembersService struct {
	client      *Client
	BaseVersion string
}

// ShareBoard Shares the board and Invites new members to collaborate on a board by sending an invitation email.
// Depending on the board's Sharing policy, there might be various scenarios where membership in the team is required in
// order to share the board with a user.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardMembersService) ShareBoard(payload ShareBoardInvitation, boardID string) (*BoardInvitationResponse, error) {
	response := &BoardInvitationResponse{}

	url := fmt.Sprintf("%s/%s/%s/%s/members", b.client.BaseURL, b.BaseVersion, EndpointBoards, boardID)

	err := b.client.Post(url, payload, response)

	return response, err
}
