package miro

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	BaseURL string
	token   string
	// HTTPClient a fine-tuned HTTP client, but you can inject your own if, for example, you wanted to lower the timeouts
	HTTPClient   *http.Client
	ctx          context.Context
	AccessToken  *AccessTokenService
	Boards       *BoardsService
	BoardMembers *BoardMembersService
	Items        *ItemsService
	AppCardItems *AppCardItemsService
	CardItems    *CardItemsService
	ShapeItems   *ShapeItemsService
}

type Field struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Context struct {
	Fields []Field `json:"fields"`
}

type ResponseError struct {
	// Status code of the error
	Status int `json:"status"`
	// Code of the error
	Code string `json:"code"`
	// Context details of the error
	Context Context `json:"context"`
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
		BaseURL:    baseURL,
		token:      token,
		HTTPClient: httpClient(),
		ctx:        context.Background(),
	}
	buildAPIMap(c)

	return c
}

func buildAPIMap(c *Client) {
	c.AccessToken = &AccessTokenService{client: c, apiVersion: "v1"}
	c.Boards = &BoardsService{client: c, apiVersion: "v2", resource: "boards"}
	c.BoardMembers = &BoardMembersService{client: c, apiVersion: "v2", resource: "boards", subResource: "members"}
	c.Items = &ItemsService{client: c, apiVersion: "v2", resource: "boards", subResource: "items"}
	c.AppCardItems = &AppCardItemsService{client: c, apiVersion: "v2", resource: "boards", subResource: "app_cards"}
	c.CardItems = &CardItemsService{client: c, apiVersion: "v2", resource: "boards", subResource: "cards"}
	c.ShapeItems = &ShapeItemsService{client: c, apiVersion: "v2", resource: "boards", subResource: "shapes"}
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
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusOK {
			return constructErrorMsg(resp)
		}
		return json.NewDecoder(resp.Body).Decode(&response)
	}
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
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusCreated {
			return constructErrorMsg(resp)
		}
		return json.NewDecoder(resp.Body).Decode(&response)
	}
}

// postNoContent Native POST function (pretending to be a DELETE method... but with query params?!)
func (c *Client) postNoContent(url string, queryParams ...Parameter) error {
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s%s", url, EncodeQueryParams(queryParams))
	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusNoContent {
			return constructErrorMsg(resp)
		}
		return nil
	}
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
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusCreated {
			return constructErrorMsg(resp)
		}
		return json.NewDecoder(resp.Body).Decode(&response)
	}
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
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusOK {
			return constructErrorMsg(resp)
		}
		return json.NewDecoder(resp.Body).Decode(&response)
	}
}

// Delete Native DELETE function
func (c *Client) Delete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	c.addHeaders(req)
	if resp, err := c.HTTPClient.Do(req); err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusNoContent {
			return constructErrorMsg(resp)
		}
		return nil
	}
}

func httpClient() *http.Client {
	transport := &http.Transport{
		// Enable keep-alive connections. By default, the http.DefaultClient does not use HTTP keep-alive, which means
		// that a new TCP connection would be established for each request.
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
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

func constructErrorMsg(resp *http.Response) error {
	respErr := &ResponseError{}
	if err := json.NewDecoder(resp.Body).Decode(respErr); err != nil {
		return err
	}

	details := make([]string, 0)
	for _, field := range respErr.Context.Fields {
		details = append(details, fmt.Sprintf("%s: %s", field.Field, field.Message))
	}
	return fmt.Errorf("unexpected status code: %d, message: %s (%s), details:\n  %s", resp.StatusCode, respErr.Message, respErr.Code, strings.Join(details, "\n  "))
}

func constructURL(urlParts ...string) (string, error) {
	for _, part := range urlParts {
		if part == "" {
			return "", errors.New("arguments cannot be empty")
		}
	}
	return strings.Join(urlParts, "/"), nil
}
