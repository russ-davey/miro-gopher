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

// Create a document item on a board by specifying the URL where the document is hosted.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) Create(boardID string, payload DocumentItemSet) (*DocumentItem, error) {
	response := &DocumentItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Upload a document item using a file from a device.
// The maximum file size supported is 28.6 MB.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) Upload(boardID, filePath string, payload UploadFileItem) (*DocumentItem, error) {
	response := &DocumentItem{}

	file, err := os.Open(filePath)
	if err != nil {
		return response, err
	}
	fileName := path.Base(file.Name())

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}

	multiParts := make(MultiParts)
	multiParts["resource"] = MultiPart{
		Reader:      io.Reader(file),
		FileName:    fileName,
		ContentType: "application/octet-stream",
	}
	multiParts["data"] = MultiPart{
		Reader:      buf,
		FileName:    fileName,
		ContentType: "application/json",
	}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.PostMultipart(c.client.ctx, url, multiParts, response)
		return response, err
	}
}

// Get information for a specific document item on a board.
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
func (c *DocumentsService) Update(boardID, itemID string, payload DocumentItemSet) (*DocumentItem, error) {
	response := &DocumentItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// UpdateFromFile update document item using a file from a device.
// Required scope: boards:write | Rate limiting: Level 2
func (c *DocumentsService) UpdateFromFile(boardID, itemID, filePath string, payload UploadFileItem) (*DocumentItem, error) {
	response := &DocumentItem{}

	file, err := os.Open(filePath)
	if err != nil {
		return response, err
	}
	fileName := path.Base(file.Name())

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}

	multiParts := make(MultiParts)
	multiParts["resource"] = MultiPart{
		Reader:      io.Reader(file),
		FileName:    fileName,
		ContentType: "application/octet-stream",
	}
	multiParts["data"] = MultiPart{
		Reader:      buf,
		FileName:    fileName,
		ContentType: "application/json",
	}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.PatchMultipart(c.client.ctx, url, multiParts, response)
		return response, err
	}
}

// Delete a document item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *DocumentsService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
