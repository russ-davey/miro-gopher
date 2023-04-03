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

// Get takes a client, API endpoint and a pointer to a struct then writes the API response data to that struct
func (c *Client) Get(url string, response interface{}, queryParams ...Arguments) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

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

func (c *Client) Post(url string, body, response interface{}) error {
	bufBody, err := bodyToBuffer(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bufBody)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

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

func (c *Client) Put(url string, body, response interface{}, queryParams ...Arguments) error {
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
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

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
