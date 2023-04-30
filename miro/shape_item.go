package miro

type ShapeItemsService struct {
	client      *Client
	APIVersion  string
	Resource    string
	SubResource string
}

// Create Adds a shape item to a board
// Required scope: boards:write | Rate limiting: Level 2
func (s *ShapeItemsService) Create(boardID string, payload SetShapeItem) (*ShapeItem, error) {
	response := &ShapeItem{}

	err := s.client.Post(s.constructURL(boardID, ""), payload, response)

	return response, err
}

// Get Retrieves information for a specific shape item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (s *ShapeItemsService) Get(boardID, itemID string) (*ShapeItem, error) {
	response := &ShapeItem{}

	err := s.client.Get(s.constructURL(boardID, itemID), response)

	return response, err
}

// Update a shape item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (s *ShapeItemsService) Update(boardID, itemID string, payload SetShapeItem) (*ShapeItem, error) {
	response := &ShapeItem{}

	err := s.client.Patch(s.constructURL(boardID, itemID), payload, response)

	return response, err
}

// Delete a shape item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (s *ShapeItemsService) Delete(boardID, itemID string) error {
	return s.client.Delete(s.constructURL(boardID, itemID))
}

func (s *ShapeItemsService) constructURL(boardID, resourceID string) string {
	return constructURL(s.client.BaseURL, s.APIVersion, s.Resource, boardID, s.SubResource, resourceID)
}
