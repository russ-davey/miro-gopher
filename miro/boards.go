package miro

import (
	"strings"
)

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
		err = b.client.Post(b.client.ctx, url, payload, response)
		return response, err
	}
}

// Get information about a board.
// Required scope: boards:read | Rate limiting: Level 1
func (b *BoardsService) Get(boardID string) (*Board, error) {
	response := &Board{}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID); err != nil {
		return response, err
	} else {
		err = b.client.Get(b.client.ctx, url, response)
		return response, err
	}
}

// GetAll boards that match the search criteria provided in the request.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: BoardSearchParams{}
func (b *BoardsService) GetAll(queryParams ...BoardSearchParams) (*ListBoards, error) {
	response := &ListBoards{client: b.client, firstResults: true}

	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource); err != nil {
		return response, err
	} else {
		if len(queryParams) > 0 {
			err = b.client.Get(b.client.ctx, url, response, parseQueryTags(queryParams[0])...)
		} else {
			err = b.client.Get(b.client.ctx, url, response)
		}

		return response, err
	}
}

// GetNext iterator to get and return a pointer to the next set of search results, starting with the first set returned by the GetAll method
func (l *ListBoards) GetNext() (*ListBoards, error) {
	response := &ListBoards{client: l.client}

	if l.firstResults {
		l.firstResults = false
		return l, nil
	}

	if l.Offset+l.Size < l.Total {
		var url string
		if idx := strings.Index(l.Links.Next, "{"); idx != -1 {
			url = l.Links.Next[:idx]
		} else {
			url = l.Links.Next
		}

		err := l.client.Get(l.client.ctx, url, response)

		return response, err
	} else {
		return response, IteratorDone
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
		err = b.client.Put(b.client.ctx, url, payload, response, Parameter{"copy_from": copyFrom})
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
		err = b.client.Patch(b.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a board.
// Required scope: boards:write | Rate limiting: Level 3
func (b *BoardsService) Delete(boardID string) error {
	if url, err := constructURL(b.client.BaseURL, b.apiVersion, b.resource, boardID); err != nil {
		return err
	} else {
		return b.client.Delete(b.client.ctx, url)
	}
}
