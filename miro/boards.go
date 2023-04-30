package miro

type BoardsService struct {
	client     *Client
	apiVersion string
	resource   string
}

// Create a board with the specified name and sharing policies.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Create(payload SetBoard) (*Board, error) {
	response := &Board{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource); err != nil {
		return response, err
	} else {
		err = b.client.Post(url, payload, response)
		return response, err
	}
}

// Get Retrieves information about a board.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardsService) Get(boardID string) (*Board, error) {
	response := &Board{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID); err != nil {
		return response, err
	} else {
		err = b.client.Get(url, response)
		return response, err
	}
}

// GetAll Retrieves a list of boards that match the search criteria provided in the request.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardSearchParams{}
func (b *BoardsService) GetAll(queryParams ...BoardSearchParams) (*ListBoards, error) {
	response := &ListBoards{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource); err != nil {
		return response, err
	} else {
		var err error
		if len(queryParams) > 0 {
			err = b.client.Get(url, response, ParseQueryTags(queryParams[0])...)
		} else {
			err = b.client.Get(url, response)
		}

		return response, err
	}
}

// Copy Creates a copy of an existing board. You can also update the name, description, sharing policy, and permissions
// policy for the new board in the request body.
// Required scope: boards:write | Rate limiting: Level 4
func (b *BoardsService) Copy(payload SetBoard, copyFrom string) (*Board, error) {
	response := &Board{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource); err != nil {
		return response, err
	} else {
		err = b.client.Put(url, payload, response, Parameter{"copy_from": copyFrom})
		return response, err
	}
}

// Update a specific board.
// Required scope: boards:write | Rate limiting: Level 2
func (b *BoardsService) Update(boardID string, payload SetBoard) (*Board, error) {
	response := &Board{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID); err != nil {
		return response, err
	} else {
		err = b.client.Patch(url, payload, response)
		return response, err
	}
}

// Delete a board.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Delete(boardID string) error {
	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID); err != nil {
		return err
	} else {
		return b.client.Delete(url)
	}
}
