package miro

type TextItemsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a text item to a board.
// Required scope: boards:write | Rate limiting: Level 2
func (t *TextItemsService) Create(boardID string, payload TextItemSet) (*TextItem, error) {
	response := &TextItem{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource); err != nil {
		return response, err
	} else {
		err = t.client.Post(t.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific text item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (t *TextItemsService) Get(boardID, itemID string) (*TextItem, error) {
	response := &TextItem{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return response, err
	} else {
		err = t.client.Get(t.client.ctx, url, response)
		return response, err
	}
}

// Update a text item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (t *TextItemsService) Update(boardID, itemID string, payload TextItemSet) (*TextItem, error) {
	response := &TextItem{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return response, err
	} else {
		err = t.client.Patch(t.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a text item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (t *TextItemsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return err
	} else {
		return t.client.Delete(t.client.ctx, url)
	}
}
