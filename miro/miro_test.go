package miro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

const (
	apiVersion = "/v2"
	testToken  = "test-token"
)

func mockMIROAPI() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()

	handler := http.NewServeMux()
	handler.Handle(apiVersion+"/", http.StripPrefix(apiVersion, mux))

	server := httptest.NewServer(handler)

	client := NewClient(testToken)

	client.BaseURL = fmt.Sprintf("%s%s", server.URL, apiVersion)
	return client, mux, server.Close
}
