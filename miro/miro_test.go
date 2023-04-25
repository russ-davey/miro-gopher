package miro

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
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

func headerExists(req *http.Request, header string) bool {
	if req.Header.Get(header) != "" {
		return true
	}
	return false
}

func TestAddHeaders(t *testing.T) {
	c := Client{}

	tests := []struct {
		method        string
		expectAuth    bool
		expectAccept  bool
		expectContent bool
	}{
		{
			method:       http.MethodGet,
			expectAuth:   true,
			expectAccept: true,
		},
		{
			method:        http.MethodPost,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:        http.MethodPut,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:        http.MethodPatch,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:     http.MethodDelete,
			expectAuth: true,
		},
	}

	Convey("Given a http method", t, func() {
		for _, test := range tests {
			Convey(fmt.Sprintf("When the addHeaders method is called with the %s method in the request", test.method), func() {
				req, _ := http.NewRequest(test.method, "http://no-where", nil)
				c.addHeaders(req)

				Convey("Then the expected headers are added to the request", func() {
					So(headerExists(req, "accept"), ShouldEqual, test.expectAccept)
					So(headerExists(req, "content-type"), ShouldEqual, test.expectContent)
					So(headerExists(req, "Authorization"), ShouldEqual, test.expectAuth)
				})
			})
		}
	})
}
