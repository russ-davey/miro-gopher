package miro

type ShapeItemsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a shape item to a board
// Required scope: boards:write | Rate limiting: Level 2
func (s *ShapeItemsService) Create(boardID string, payload SetShapeItem) (*ShapeItem, error) {
	response := &ShapeItem{}

	if url, err := constructURL(s.client.BaseURL, s.apiVersion, s.resource, boardID, s.subResource); err != nil {
		return response, err
	} else {
		err = s.client.Post(url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific shape item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (s *ShapeItemsService) Get(boardID, itemID string) (*ShapeItem, error) {
	response := &ShapeItem{}

	if url, err := constructURL(s.client.BaseURL, s.apiVersion, s.resource, boardID, s.subResource, itemID); err != nil {
		return response, err
	} else {
		err = s.client.Get(url, response)
		return response, err
	}
}

// Update a shape item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (s *ShapeItemsService) Update(boardID, itemID string, payload SetShapeItem) (*ShapeItem, error) {
	response := &ShapeItem{}

	if url, err := constructURL(s.client.BaseURL, s.apiVersion, s.resource, boardID, s.subResource, itemID); err != nil {
		return response, err
	} else {
		err = s.client.Patch(url, payload, response)
		return response, err
	}
}

// Delete a shape item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (s *ShapeItemsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(s.client.BaseURL, s.apiVersion, s.resource, boardID, s.subResource, itemID); err != nil {
		return err
	} else {
		return s.client.Delete(url)
	}
}
