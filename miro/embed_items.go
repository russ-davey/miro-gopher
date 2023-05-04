package miro

type EmbedItemsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds an embed item to a board
// Required scope: boards:write | Rate limiting: Level 2
func (c *EmbedItemsService) Create(boardID string, payload SetEmbedItem) (*EmbedItem, error) {
	response := &EmbedItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific embed item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (c *EmbedItemsService) Get(boardID, itemID string) (*EmbedItem, error) {
	response := &EmbedItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// Update a embed item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (c *EmbedItemsService) Update(boardID, itemID string, payload SetEmbedItem) (*EmbedItem, error) {
	response := &EmbedItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a embed item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *EmbedItemsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
