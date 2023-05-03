package miro

type ConnectorsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a connector to a board.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ConnectorsService) Create(boardID string, payload SetConnector) (*Connector, error) {
	response := &Connector{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Get Retrieves information for a specific connector on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (c *ConnectorsService) Get(boardID, itemID string) (*Connector, error) {
	response := &Connector{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// GetAll Retrieves a list of connectors for a specific board.
//
// This method returns results using a cursor-based approach. A cursor-paginated method returns a portion of the total
// set of results based on the limit specified and a cursor that points to the next portion of the results.
// To retrieve the next portion of the collection, on your next call to the same method, set the cursor parameter equal
// to the cursor value you received in the response of the previous request. For example, if you set the limit query
// parameter to 10 and the board contains 20 objects, the first call will return information about the first 10 objects
// in the response along with a cursor parameter and value. In this example, let's say the cursor parameter value
// returned in the response is foo. If you want to retrieve the next set of objects, on your next call to the same method,
// set the cursor parameter value to foo.
// Required scope: boards:read | Rate limiting: Level 2
// Search query params: ConnectorsSearchParams{}
func (c *ConnectorsService) GetAll(boardID string, queryParams ...ConnectorSearchParams) (*ListConnectors, error) {
	response := &ListConnectors{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		var err error
		if len(queryParams) > 0 {
			err = c.client.Get(c.client.ctx, url, response, ParseQueryTags(queryParams[0])...)
		} else {
			err = c.client.Get(c.client.ctx, url, response)
		}

		return response, err
	}
}

// Update a connector on a board based on the data and style properties provided in the request body.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ConnectorsService) Update(boardID, itemID string, payload SetConnector) (*Connector, error) {
	response := &Connector{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Delete the specified connector from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *ConnectorsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
