package miro

type StickyNotesService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create a sticky note item on a board
// Required scope: boards:write | Rate limiting: Level 2
func (c *StickyNotesService) Create(boardID string, payload StickyNoteSet) (*StickyNote, error) {
	response := &StickyNote{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Get information for a specific sticky note item on a board
// Required scope: boards:read | Rate limiting: Level 1
func (c *StickyNotesService) Get(boardID, itemID string) (*StickyNote, error) {
	response := &StickyNote{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// Update a sticky note item on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (c *StickyNotesService) Update(boardID, itemID string, payload StickyNoteSet) (*StickyNote, error) {
	response := &StickyNote{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a sticky note item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *StickyNotesService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
