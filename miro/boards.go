package miro

import (
	"fmt"
)

const (
	// EndpointBoards /boards endpoint
	EndpointBoards = "boards"

	// QueryParamCopyFrom Unique identifier (ID) of the board that you want to copy (required).
	QueryParamCopyFrom = "copy_from"
)

type BoardsService struct {
	client      *Client
	BaseVersion string
}

// Create a board with the specified name and sharing policies.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Create(body CreateBoard) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards)
	err := b.client.Post(url, body, response)

	return response, err
}

// Get Retrieves information about a board.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardsService) Get(id string) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards, id)
	err := b.client.Get(url, response)

	return response, err
}

// GetAll Retrieves a list of boards that match the search criteria provided in the request.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardSearchParams{}
func (b *BoardsService) GetAll(queryParams ...BoardSearchParams) (*ListBoards, error) {
	response := &ListBoards{}

	url := fmt.Sprintf("%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards)

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

	url := fmt.Sprintf("%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards)
	err := b.client.Put(url, body, response, Parameter{
		QueryParamCopyFrom: copyFrom,
	})

	return response, err
}

// Update Updates a specific board.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardsService) Update(body CreateBoard, id string) (*Board, error) {
	response := &Board{}

	url := fmt.Sprintf("%s/%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards, id)
	err := b.client.Patch(url, body, response)

	return response, err
}

// Delete Deletes a board.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Delete(id string) error {
	return b.client.Delete(fmt.Sprintf("%s/%s/%s/%s", b.client.BaseURL, b.BaseVersion, EndpointBoards, id))
}
