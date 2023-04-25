package miro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	token   string
	ctx     context.Context
	Boards  *BoardsService
}

type ResponseError struct {
	// Status code of the error
	Status int `json:"status"`
	// Code of the error
	Code string `json:"code"`
	// Description of the error
	Message string `json:"message"`
	// Type of the error
	Type string `json:"type"`
}

func NewClient(token string) *Client {
	var baseURL string
	if mockServer := os.Getenv("MIRO_MOCK_SERVER"); mockServer != "" {
		baseURL = mockServer
	} else {
		baseURL = "https://api.miro.com/v2"
	}

	c := &Client{
		BaseURL: baseURL,
		token:   token,
		ctx:     context.Background(),
	}
	c.Boards = &BoardsService{client: c}

	return c
}

// Get Native GET function
func (c *Client) Get(url string, response interface{}, queryParams ...Parameter) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		respErr := &ResponseError{}
		if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Post Native POST function
func (c *Client) Post(url string, body, response interface{}) error {
	bufBody, err := bodyToBuffer(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		respErr := &ResponseError{}
		if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Put Native PUT function
func (c *Client) Put(url string, body, response interface{}, queryParams ...Parameter) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	bufBody, err := bodyToBuffer(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		respErr := &ResponseError{}
		if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Patch Native PATCH function
func (c *Client) Patch(url string, body, response interface{}) error {
	bufBody, err := bodyToBuffer(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		respErr := &ResponseError{}
		if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Delete Native DELETE function
func (c *Client) Delete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		respErr := &ResponseError{}
		if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
	}

	return nil
}

func (c *Client) addHeaders(r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		r.Header.Add("accept", "application/json")
	case http.MethodDelete:
	default:
		r.Header.Add("accept", "application/json")
		r.Header.Add("content-type", "application/json")
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

}

func bodyToBuffer(body interface{}) (io.ReadWriter, error) {
	var bufBody io.ReadWriter
	if body != nil {
		if jsonBody, err := json.Marshal(body); err != nil {
			return nil, err
		} else {
			bufBody = bytes.NewBuffer(jsonBody)
		}
	}
	return bufBody, nil
}
