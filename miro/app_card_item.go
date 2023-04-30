package miro

import "fmt"

type AppCardItemsService struct {
	client      *Client
	APIVersion  string
	SubResource string
}

// Create Adds an app card item to a board.
// Required scope: boards:write | Rate limiting: Level 2
func (a *AppCardItemsService) Create(boardID string, payload SetAppCardItem) (*AppCardItem, error) {
	response := &AppCardItem{}

	err := a.client.Post(a.constructURL(boardID, ""), payload, response)

	return response, err
}

// Get Retrieves information for a specific app card item on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (a *AppCardItemsService) Get(boardID, itemID string) (*AppCardItem, error) {
	response := &AppCardItem{}

	err := a.client.Get(a.constructURL(boardID, itemID), response)

	return response, err
}

// Update an app card item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (a *AppCardItemsService) Update(boardID, itemID string, payload SetAppCardItem) (*AppCardItem, error) {
	response := &AppCardItem{}

	err := a.client.Patch(a.constructURL(boardID, itemID), payload, response)

	return response, err
}

// Delete an app card item from a board.
// Required scope: boards:write | Rate limiting: Level 3
func (a *AppCardItemsService) Delete(boardID, itemID string) error {
	return a.client.Delete(a.constructURL(boardID, itemID))
}

func (a *AppCardItemsService) constructURL(boardID, resourceID string) string {
	if resourceID != "" {
		return fmt.Sprintf("%s/%s/%s/%s/%s/%s", a.client.BaseURL, a.APIVersion, EndpointBoards, boardID, a.SubResource, resourceID)
	} else {
		return fmt.Sprintf("%s/%s/%s/%s/%s", a.client.BaseURL, a.APIVersion, EndpointBoards, boardID, a.SubResource)
	}
}
