package miro

import "fmt"

type CardItemsService struct {
	client      *Client
	APIVersion  string
	SubResource string
}

// Create Adds a card item to a board
// Required scope: boards:write | Rate limiting: Level 2
func (c *CardItemsService) Create(boardID string, payload CreateCardItem) (*CardItem, error) {
	response := &CardItem{}

	err := c.client.Post(c.constructURL(boardID, ""), payload, response)

	return response, err
}

// Get Retrieves information for a specific card item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (c *CardItemsService) Get(boardID, itemID string) (*CardItem, error) {
	response := &CardItem{}

	err := c.client.Get(c.constructURL(boardID, itemID), response)

	return response, err
}

// Update a card item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (c *CardItemsService) Update(boardID, itemID string, itemUpdate CreateCardItem) (*CardItem, error) {
	response := &CardItem{}

	err := c.client.Patch(c.constructURL(boardID, itemID), itemUpdate, response)

	return response, err
}

// Delete a card item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *CardItemsService) Delete(boardID, itemID string) error {
	return c.client.Delete(c.constructURL(boardID, itemID))
}

func (c *CardItemsService) constructURL(boardID, resourceID string) string {
	if resourceID != "" {
		return fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.client.BaseURL, c.APIVersion, EndpointBoards, boardID, c.SubResource, resourceID)
	} else {
		return fmt.Sprintf("%s/%s/%s/%s/%s", c.client.BaseURL, c.APIVersion, EndpointBoards, boardID, c.SubResource)
	}
}
