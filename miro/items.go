package miro

type ItemsService struct {
	client      *Client
	APIVersion  string
	Resource    string
	SubResource string
}

// GetAll Retrieves a list of items for a specific board. You can retrieve all items on the board, a list of child items
// inside a parent item, or a list of specific types of items by specifying URL query parameter values.
//
// This method returns results using a cursor-based approach. A cursor-paginated method returns a portion of the total set
// of results based on the limit specified and a cursor that points to the next portion of the results.
// To retrieve the next portion of the collection, on your next call to the same method, set the cursor parameter equal
// to the cursor value you received in the response of the previous request. For example, if you set the limit query
// parameter to 10 and the board contains 20 objects, the first call will return information about the first 10 objects
// in the response along with a cursor parameter and value. In this example, let's say the cursor parameter value returned
// in the response is foo. If you want to retrieve the next set of objects, on your next call to the same method, set the
// cursor parameter value to foo.
// Required scope: boards:read | Rate limiting: Level 2
func (i *ItemsService) GetAll(boardID string, queryParams ...ItemSearchParams) (*ListItems, error) {
	response := &ListItems{}

	url := i.constructURL(boardID, "")

	var err error
	if len(queryParams) > 0 {
		err = i.client.Get(url, response, ParseQueryTags(queryParams[0])...)
	} else {
		err = i.client.Get(url, response)
	}

	return response, err
}

// Get Retrieves information for a specific item on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (i *ItemsService) Get(boardID, itemID string) (*Item, error) {
	response := &Item{}

	err := i.client.Get(i.constructURL(boardID, itemID), response)

	return response, err
}

// Update item position or parent
// Required scope: boards:write | Rate limiting: Level 2
func (i *ItemsService) Update(boardID, itemID string, itemUpdate ItemUpdate) (*Item, error) {
	response := &Item{}

	err := i.client.Patch(i.constructURL(boardID, itemID), itemUpdate, response)

	return response, err
}

// Delete an item from a board.
// Required scope: boards:write | Rate limiting: Level 3
func (i *ItemsService) Delete(boardID, itemID string) error {
	return i.client.Delete(i.constructURL(boardID, itemID))
}

func (i *ItemsService) constructURL(boardID, resourceID string) string {
	return constructURL(i.client.BaseURL, i.APIVersion, i.Resource, boardID, i.SubResource, resourceID)
}
