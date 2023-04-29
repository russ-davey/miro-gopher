package miro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	BaseURL      string
	token        string
	ctx          context.Context
	AccessToken  *AccessTokenService
	Boards       *BoardsService
	BoardMembers *BoardMembersService
	Items        *ItemsService
	AppCardItems *AppCardItemsService
	CardItems    *CardItemsService
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
		baseURL = "https://api.miro.com"
	}

	c := &Client{
		BaseURL: baseURL,
		token:   token,
		ctx:     context.Background(),
	}
	buildAPIMap(c)

	return c
}

func buildAPIMap(c *Client) {
	c.AccessToken = &AccessTokenService{client: c, APIVersion: "v1"}
	c.Boards = &BoardsService{client: c, APIVersion: "v2"}
	c.BoardMembers = &BoardMembersService{client: c, APIVersion: "v2", SubResource: "members"}
	c.Items = &ItemsService{client: c, APIVersion: "v2", SubResource: "items"}
	c.AppCardItems = &AppCardItemsService{client: c, APIVersion: "v2", SubResource: "app_cards"}
	c.CardItems = &CardItemsService{client: c, APIVersion: "v2", SubResource: "cards"}
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
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return constructErrorMsg(res)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Post Native POST function
func (c *Client) Post(url string, payload, response interface{}) error {
	bufBody, err := payloadToBuffer(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return constructErrorMsg(res)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// PostNoContent Native POST function (pretending to be a DELETE method... but with query params?!)
func (c *Client) PostNoContent(url string, queryParams ...Parameter) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return constructErrorMsg(res)
	}
	return nil
}

// Put Native PUT function
func (c *Client) Put(url string, payload, response interface{}, queryParams ...Parameter) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	bufBody, err := payloadToBuffer(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return constructErrorMsg(res)
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

// Patch Native PATCH function
func (c *Client) Patch(url string, payload, response interface{}) error {
	bufBody, err := payloadToBuffer(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bufBody)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return constructErrorMsg(res)
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
	res, err := httpClient().Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return constructErrorMsg(res)
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

func payloadToBuffer(body interface{}) (io.ReadWriter, error) {
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

func constructErrorMsg(res *http.Response) error {
	respErr := &ResponseError{}
	if err := json.NewDecoder(res.Body).Decode(respErr); err != nil {
		return err
	}
	return fmt.Errorf("unexpected status code: %d, message: %s (%s)", res.StatusCode, respErr.Message, respErr.Code)
}

func httpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
