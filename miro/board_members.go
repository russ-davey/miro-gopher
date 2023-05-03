package miro

type BoardMembersService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// ShareBoard Shares the board and Invites new members to collaborate on a board by sending an invitation email.
// Depending on the board's Sharing policy, there might be various scenarios where membership in the team is required in
// order to share the board with a user.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardMembersService) ShareBoard(boardID string, payload ShareBoardInvitation) (*BoardInvitationResponse, error) {
	response := &BoardInvitationResponse{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID, b.subResource); err != nil {
		return response, err
	} else {
		err = b.client.Post(b.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a board member.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardMembersService) Get(boardID, itemID string) (*BoardMember, error) {
	response := &BoardMember{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID, b.subResource, itemID); err != nil {
		return response, err
	} else {
		err = b.client.Get(b.client.ctx, url, response)
		return response, err
	}
}

// GetAll Retrieves a pageable list of members for a board.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardMemberSearchParams{}
func (b *BoardMembersService) GetAll(boardID string, queryParams ...BoardMemberSearchParams) (*ListBoardMembers, error) {
	response := &ListBoardMembers{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID, b.subResource); err != nil {
		return response, err
	} else {
		var err error
		if len(queryParams) > 0 {
			err = b.client.Get(b.client.ctx, url, response, ParseQueryTags(queryParams[0])...)
		} else {
			err = b.client.Get(b.client.ctx, url, response)
		}

		return response, err
	}
}

// Update the role of a board member.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardMembersService) Update(boardID, itemID string, role Role) (*BoardMember, error) {
	response := &BoardMember{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID, b.subResource, itemID); err != nil {
		return response, err
	} else {
		err = b.client.Patch(b.client.ctx, url, RoleUpdate{Role: role}, response)
		return response, err
	}
}

// Delete Removes a board member from a board.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardMembersService) Delete(boardID, itemID string) error {
	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID, b.subResource, itemID); err != nil {
		return err
	} else {
		return b.client.Delete(b.client.ctx, url)
	}
}
