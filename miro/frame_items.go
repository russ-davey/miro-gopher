package miro

type FramesService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a frame to a board.
// Required scope: boards:write | Rate limiting: Level 2
func (f *FramesService) Create(boardID string, payload SetFrameItem) (*FrameItem, error) {
	response := &FrameItem{}

	addPayloadDefaults(&payload)

	if url, err := constructURL(f.client.BaseURL, f.apiVersion, f.resource, boardID, f.subResource); err != nil {
		return response, err
	} else {
		err = f.client.Post(f.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific frame on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (f *FramesService) Get(boardID, itemID string) (*FrameItem, error) {
	response := &FrameItem{}

	if url, err := constructURL(f.client.BaseURL, f.apiVersion, f.resource, boardID, f.subResource, itemID); err != nil {
		return response, err
	} else {
		err = f.client.Get(f.client.ctx, url, response)
		return response, err
	}
}

// GetItems Retrieves a list of items within a specific frame. A frame is a parent item and all items within a frame are child items.
// This method returns results using a cursor-based approach. A cursor-paginated method returns a portion of the total
// set of results based on the limit specified and a cursor that points to the next portion of the results.
// To retrieve the next portion of the collection, on your next call to the same method, set the cursor parameter equal
// to the cursor value you received in the response of the previous request. For example, if you set the limit query
// parameter to 10 and the board contains 20 objects, the first call will return information about the first 10 objects
// in the response along with a cursor parameter and value. In this example, let's say the cursor parameter value
// returned in the response is foo. If you want to retrieve the next set of objects, on your next call to the same method,
// set the cursor parameter value to foo.
// Required scope: boards:read | Rate limiting: Level 2
// Search query params: ItemSearchParams{}
func (f *FramesService) GetItems(boardID, frameID string, queryParams ...ItemSearchParams) (*ListItems, error) {
	response := &ListItems{}

	if url, err := constructURL(f.client.BaseURL, f.apiVersion, f.resource, boardID, "items"); err != nil {
		return response, err
	} else {
		var err error
		if len(queryParams) > 0 {
			searchParams := ParseQueryTags(queryParams[0])
			searchParams = append(searchParams, Parameter{"parent_item_id": frameID})
			err = f.client.Get(f.client.ctx, url, response, searchParams...)
		} else {
			err = f.client.Get(f.client.ctx, url, response, Parameter{"parent_item_id": frameID})
		}

		return response, err
	}
}

// Update a frame on a board based on the data, style, or geometry properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (f *FramesService) Update(boardID, itemID string, payload SetFrameItem) (*FrameItem, error) {
	response := &FrameItem{}

	addPayloadDefaults(&payload)

	if url, err := constructURL(f.client.BaseURL, f.apiVersion, f.resource, boardID, f.subResource, itemID); err != nil {
		return response, err
	} else {
		err = f.client.Patch(f.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete the specified frame from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (f *FramesService) Delete(boardID, itemID string) error {
	if url, err := constructURL(f.client.BaseURL, f.apiVersion, f.resource, boardID, f.subResource, itemID); err != nil {
		return err
	} else {
		return f.client.Delete(f.client.ctx, url)
	}
}

func addPayloadDefaults(payload *SetFrameItem) {
	if payload.Data.Type == "" {
		payload.Data.Type = TypeFreeform
	}
	if payload.Data.Format == "" {
		payload.Data.Format = FormatCustom
	}
}
