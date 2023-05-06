package miro

type TagsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// GetTags Retrieves all the items that have the specified tag.
// Required scope: boards:read | Rate limiting: Level 1
// Search query params: TagSearchParams{}
func (t *TagsService) GetTags(boardID, tagID string, queryParams ...TagSearchParams) (*ListItems, error) {
	response := &ListItems{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, "items"); err != nil {
		return response, err
	} else {
		var searchParams []Parameter
		if len(queryParams) > 0 {
			searchParams = parseQueryTags(queryParams[0])
		}
		searchParams = append(searchParams, Parameter{"tag_id": tagID})

		err = t.client.Get(t.client.ctx, url, response, searchParams...)
		return response, err
	}
}

// Attach an existing tag to the specified item. Card and sticky note items can have up to 8 tags.
// Required scope: boards:write | Rate limiting: Level 1
func (t *TagsService) Attach(boardID, itemID, tagID string) error {
	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, "items", itemID); err != nil {
		return err
	} else {
		return t.client.postNoContent(t.client.ctx, url, Parameter{"tag_id": tagID})
	}
}

// Detach removes the specified tag from the specified item. The tag still exists on the board.
// Required scope: boards:write | Rate limiting: Level 1
func (t *TagsService) Detach(boardID, itemID, tagID string) error {
	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, "items", itemID); err != nil {
		return err
	} else {
		return t.client.Delete(t.client.ctx, url, Parameter{"tag_id": tagID})
	}
}

func (t *TagsService) GetTagsFromItem(boardID, itemID string) (*ListTags, error) {
	response := &ListTags{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, "items", itemID, "tags"); err != nil {
		return response, err
	} else {
		err := t.client.Get(t.client.ctx, url, response)

		return response, err
	}
}

// Create a tag on a board.
// Required scope: boards:write | Rate limiting: Level 1
func (t *TagsService) Create(boardID string, payload TagSet) (*Tag, error) {
	response := &Tag{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource); err != nil {
		return response, err
	} else {
		err = t.client.Post(t.client.ctx, url, payload, response)
		return response, err
	}
}

func (t *TagsService) GetTagsFromBoard(boardID string, queryParams ...TagSearchParams) (*ListBoardTags, error) {
	response := &ListBoardTags{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, "tags"); err != nil {
		return response, err
	} else {
		if len(queryParams) > 0 {
			err = t.client.Get(t.client.ctx, url, response, parseQueryTags(queryParams[0])...)
		} else {
			err = t.client.Get(t.client.ctx, url, response)
		}

		return response, err
	}
}

// Get Retrieves information for a specific tag.
// Required scope: boards:read | Rate limiting: Level 1
func (t *TagsService) Get(boardID, itemID string) (*Tag, error) {
	response := &Tag{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return response, err
	} else {
		err = t.client.Get(t.client.ctx, url, response)
		return response, err
	}
}

// Update a tag based on the data properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 1
func (t *TagsService) Update(boardID, itemID string, payload TagSet) (*Tag, error) {
	response := &Tag{}

	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return response, err
	} else {
		err = t.client.Patch(t.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete a specified tag from the board. The tag is also removed from all cards and sticky notes on the board.
// Required scope: boards:write | Rate limiting: Level 1
func (t *TagsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(t.client.BaseURL, t.apiVersion, t.resource, boardID, t.subResource, itemID); err != nil {
		return err
	} else {
		return t.client.Delete(t.client.ctx, url)
	}
}
