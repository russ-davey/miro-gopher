package miro

import (
	"fmt"
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

	err := b.client.Post(b.constructURL(boardID), payload, response)

	return response, err
}

// GetAll Retrieves a pageable list of members for a board.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardMemberSearchParams{}
func (b *BoardMembersService) GetAll(boardID string, queryParams ...BoardMemberSearchParams) (*ListBoardMembersResponse, error) {
	response := &ListBoardMembersResponse{}

	url := b.constructURL(boardID)

	var err error
	if len(queryParams) > 0 {
		err = b.client.Get(url, response, ParseQueryTags(queryParams[0])...)
	} else {
		err = b.client.Get(url, response)
	}

	return response, err
}

func (b *BoardMembersService) constructURL(boardID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/members", b.client.BaseURL, b.BaseVersion, EndpointBoards, boardID)
}
