package miro

import (
	"fmt"
)

const (
	EndpointBoards = "boards"

	// QueryParamCopyFrom Unique identifier (ID) of the board that you want to copy (required).
	QueryParamCopyFrom = "copy_from"

	// SortDefault If team_id is present, last_created. Otherwise, last_opened
	SortDefault = "default"
	// SortLastModified sort by the date and time when the board was last modified
	SortLastModified = "last_modified"
	// SortLastOpened sort by the date and time when the board was last opened
	SortLastOpened = "last_opened"
	// SortLastCreated sort by the date and time when the board was created
	SortLastCreated = "last_created"
	// SortAlphabetically sort by the board name (alphabetically)
	SortAlphabetically = "alphabetically"

	CopyAccessAnyone      = "anyone"
	CopyAccessTeamMembers = "team_members"
	CopyAccessTeamEditors = "team_editors"
	CopyAccessBoardOwner  = "board_owner"

	AccessTeamMemberWithEditingRights = "team_members_with_editing_rights"
	AccessBoardOwnersAndCoOwners      = "board_owners_and_coowners"

	CollabToolsStartAccessAllEditors = "all_editors"

	AccessPrivate = "private"
	AccessView    = "view"
	AccessComment = "comment"
	AccessEdit    = "edit"

	InviteAccessViewer    = "viewer"
	InviteAccessCommenter = "commenter"
	InviteAccessEditor    = "editor"
	InviteAccessNoAccess  = "no_access"
)

type BoardsService struct {
	client *Client
}

type BoardSearchParams struct {
	// TeamID The team_id for which you want to retrieve the list of boards. If this parameter is sent in the request,
	// the query and owner parameters are ignored.
	TeamID string `query:"team_id,omitempty"`
	// Query Retrieves a list of boards that contain the query string provided in the board content or board name.
	// For example, if you want to retrieve a list of boards that contain the word beta within the board itself (board content),
	// add beta as the query parameter value. You can use the query parameter with the owner parameter to narrow down the board search results.
	Query string `query:"query,omitempty"`
	// Owner Retrieves a list of boards that belong to a specific owner ID. You must pass the owner ID (for example,
	//3074457353169356300), not the owner name. You can use the 'owner' parameter with the query parameter to narrow
	//down the board search results. Note that if you pass the team_id in the same request, the owner parameter is ignored.
	Owner string `query:"owner,omitempty"`
	// Limit The maximum number of boards to retrieve.
	// Default: 20
	Limit string `query:"limit,omitempty"`
	// The (zero-based) offset of the first item in the collection to return.
	// Default: 0.
	Offset string `query:"offset,omitempty"`
	// Sort The order in which you want to view the result set.
	// Options last_created and alphabetically are applicable only when you search for boards by team.
	Sort string `query:"sort,omitempty"`
}

// Create a board with the specified name and sharing policies.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Create(body CreateBoard) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s", b.client.BaseURL, EndpointBoards)
	err := b.client.Post(url, body, response)

	return response, err
}

// Get Retrieves information about a board.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardsService) Get(id string) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s/%s", b.client.BaseURL, EndpointBoards, id)
	err := b.client.Get(url, response)

	return response, err
}

// GetAll Retrieves a list of boards that match the search criteria provided in the request.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardSearchParams{}
func (b *BoardsService) GetAll(queryParams ...BoardSearchParams) (*ListBoards, error) {
	response := &ListBoards{}

	url := fmt.Sprintf("%s/%s", b.client.BaseURL, EndpointBoards)

	var err error
	if len(queryParams) > 0 {
		err = b.client.Get(url, response, ParseQueryTags(queryParams[0])...)
	} else {
		err = b.client.Get(url, response)
	}

	return response, err
}

// Copy Creates a copy of an existing board. You can also update the name, description, sharing policy, and permissions
// policy for the new board in the request body.
// Required scope: boards:write | Rate limiting: Level 4
func (b *BoardsService) Copy(body CreateBoard, copyFrom string) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s", b.client.BaseURL, EndpointBoards)
	err := b.client.Put(url, body, response, Parameter{
		QueryParamCopyFrom: copyFrom,
	})

	return response, err
}

// Update Updates a specific board.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardsService) Update(body CreateBoard, id string) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s/%s", b.client.BaseURL, EndpointBoards, id)
	err := b.client.Patch(url, body, response)

	return response, err
}
