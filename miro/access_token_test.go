package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client, _, mux, closeAPIServer := mockMIROAPI("v1", "", "", "")
	defer closeAPIServer()

	expectedResults := &AccessToken{}
	responseData := constructResponseAndResults("access_token_get.json", expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given no arguments", t, func() {
		Convey("When the AccessToken Get function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v1/%s", EndpointOAUTHToken), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.AccessToken.Get()

			Convey("Then the access token data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v1/%s", EndpointOAUTHToken))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestRevokeAccessToken(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v1", EndpointOAUTH, "revoke", "")
	defer closeAPIServer()

	Convey("Given an access token", t, func() {
		Convey("When the AccessToken Revoke function is called", func() {
			var receivedRequest *http.Request

			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.AccessToken.Revoke(testToken)

			Convey("Then the access token is revoked", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPost)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Query().Get("access_token"), ShouldEqual, testToken)
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v1/%s/revoke", EndpointOAUTH))
				})
			})
		})
	})
}
