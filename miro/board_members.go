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

	err := b.client.Post(b.constructURL(boardID, ""), payload, response)

	return response, err
}

// Get Retrieves information for a board member.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardMembersService) Get(boardID, boardMemberID string) (*BoardMember, error) {
	response := &BoardMember{}

	err := b.client.Get(b.constructURL(boardID, boardMemberID), response)

	return response, err
}

// GetAll Retrieves a pageable list of members for a board.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardMemberSearchParams{}
func (b *BoardMembersService) GetAll(boardID string, queryParams ...BoardMemberSearchParams) (*ListBoardMembersResponse, error) {
	response := &ListBoardMembersResponse{}

	url := b.constructURL(boardID, "")

	var err error
	if len(queryParams) > 0 {
		err = b.client.Get(url, response, ParseQueryTags(queryParams[0])...)
	} else {
		err = b.client.Get(url, response)
	}

	return response, err
}

// Update Updates the role of a board member.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardMembersService) Update(boardID, boardMemberID string, role Role) (*BoardMember, error) {
	response := &BoardMember{}

	err := b.client.Patch(b.constructURL(boardID, boardMemberID), RoleUpdate{Role: role}, response)

	return response, err
}

// Delete Removes a board member from a board.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardMembersService) Delete(boardID, boardMemberID string) error {
	return b.client.Delete(b.constructURL(boardID, boardMemberID))
}

func (b *BoardMembersService) constructURL(boardID, boardMemberID string) string {
	if boardMemberID != "" {
		return fmt.Sprintf("%s/%s/%s/%s/members/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards, boardID, boardMemberID)
	} else {
		return fmt.Sprintf("%s/%s/%s/%s/members", b.client.BaseURL, b.BaseVersion, EndpointBoards, boardID)
	}
}
