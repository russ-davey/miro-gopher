package miro

type AppCardItemsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds an app card item to a board.
// Required scope: boards:write | Rate limiting: Level 2
func (a *AppCardItemsService) Create(boardID string, payload SetAppCardItem) (*AppCardItem, error) {
	response := &AppCardItem{}

	if url, err := constructURL(a.client.BaseURL, a.apiVersion, a.resource, boardID, a.subResource); err != nil {
		return response, err
	} else {
		err = a.client.Post(a.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific app card item on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (a *AppCardItemsService) Get(boardID, itemID string) (*AppCardItem, error) {
	response := &AppCardItem{}

	if url, err := constructURL(a.client.BaseURL, a.apiVersion, a.resource, boardID, a.subResource, itemID); err != nil {
		return response, err
	} else {
		err = a.client.Get(a.client.ctx, url, response)
		return response, err
	}
}

// Update an app card item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (a *AppCardItemsService) Update(boardID, itemID string, payload SetAppCardItem) (*AppCardItem, error) {
	response := &AppCardItem{}

	if url, err := constructURL(a.client.BaseURL, a.apiVersion, a.resource, boardID, a.subResource, itemID); err != nil {
		return response, err
	} else {
		err = a.client.Patch(a.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete an app card item from a board.
// Required scope: boards:write | Rate limiting: Level 3
func (a *AppCardItemsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(a.client.BaseURL, a.apiVersion, a.resource, boardID, a.subResource, itemID); err != nil {
		return err
	} else {
		return a.client.Delete(a.client.ctx, url)
	}
}
