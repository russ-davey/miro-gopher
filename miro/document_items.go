package miro

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path"
)

type DocumentsService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a document item to a board by specifying the URL where the document is hosted.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) Create(boardID string, payload SetDocumentItem) (*DocumentItem, error) {
	response := &DocumentItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Upload Creates a document item using file from device.
// The maximum file size supported is 28.6 MB.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) Upload(boardID, filePath string, payload UploadFileItem) (*DocumentItem, error) {
	response := &DocumentItem{}

	file, err := os.Open(filePath)
	if err != nil {
		return response, err
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}

	reader := make(map[string]io.Reader)
	reader["resource"] = io.Reader(file)
	reader["data"] = buf

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.PostMultipart(c.client.ctx, url, path.Base(file.Name()), reader, response)
		return response, err
	}
}

// Get Retrieves information for a specific document item on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (c *DocumentsService) Get(boardID, itemID string) (*DocumentItem, error) {
	response := &DocumentItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// Update a document item on a board.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) Update(boardID, itemID string, payload SetDocumentItem) (*DocumentItem, error) {
	response := &DocumentItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

//// UpdateFromFile a document item on a board. Update document item using file from device.
//// Required scope: boards:write | Rate limiting: Level 2
//func (c *DocumentsService) UpdateFromFile(boardID, itemID string, payload SetDocumentItem) (*DocumentItem, error) {
//	response := &DocumentItem{}
//
//	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
//		return response, err
//	} else {
//		err = c.client.Patch(c.client.ctx, url, payload, response)
//		return response, err
//	}
//}

// Delete a document item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *DocumentsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
