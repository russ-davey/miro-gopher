package miro

type CardItemsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create a card item on a board
// Required scope: boards:write | Rate limiting: Level 2
func (c *CardItemsService) Create(boardID string, payload SetCardItem) (*CardItem, error) {
	response := &CardItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Get information for a specific card item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (c *CardItemsService) Get(boardID, itemID string) (*CardItem, error) {
	response := &CardItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// Update a card item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (c *CardItemsService) Update(boardID, itemID string, payload SetCardItem) (*CardItem, error) {
	response := &CardItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a card item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *CardItemsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
