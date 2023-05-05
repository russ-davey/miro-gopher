package miro

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path"
)

type ImagesService struct {
	client      *Client
	apiVersion  string
	resource    string
	subResource string
}

// Create Adds a image item to a board by specifying the URL where the image is hosted.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ImagesService) Create(boardID string, payload ImageItemSet) (*ImageItem, error) {
	response := &ImageItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource); err != nil {
		return response, err
	} else {
		err = c.client.Post(c.client.ctx, url, payload, response)
		return response, err
	}
}

// Upload Creates an image item using file from device.
// The maximum file size supported is 28.6 MB.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ImagesService) Upload(boardID, filePath string, payload UploadFileItem) (*ImageItem, error) {
	response := &ImageItem{}

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

// Get Retrieves information for a specific image item on a board.
// Required scope: boards:read | Rate limiting: Level 1
func (c *ImagesService) Get(boardID, itemID string) (*ImageItem, error) {
	response := &ImageItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Get(c.client.ctx, url, response)
		return response, err
	}
}

// Update a image item on a board.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ImagesService) Update(boardID, itemID string, payload ImageItemSet) (*ImageItem, error) {
	response := &ImageItem{}

	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return response, err
	} else {
		err = c.client.Patch(c.client.ctx, url, payload, response)
		return response, err
	}
}

// UpdateFromFile update image item using file from device.
// Required scope: boards:write | Rate limiting: Level 2
func (c *ImagesService) UpdateFromFile(boardID, itemID, filePath string, payload UploadFileItem) (*ImageItem, error) {
	response := &ImageItem{}

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

// Delete a image item from the board.
// Required scope: boards:write | Rate limiting: Level 3
func (c *ImagesService) Delete(boardID, itemID string) error {
	if url, err := constructURL(c.client.BaseURL, c.apiVersion, c.resource, boardID, c.subResource, itemID); err != nil {
		return err
	} else {
		return c.client.Delete(c.client.ctx, url)
	}
}
