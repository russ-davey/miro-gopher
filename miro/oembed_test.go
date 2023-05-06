package miro

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestGetOEmbed(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v1", "oembed", "", "")

	defer closeAPIServer()

	expectedResults := &OEmbed{}
	responseData := constructResponseAndResults("oembed_get.json", expectedResults)

	Convey("Given a URL", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.OEmbed.Get("http://testing")

			Convey("Then the oEmbed information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("url"), ShouldEqual, "http://testing")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, testResourcePath)
				})
			})
		})
	})
}

func TestGetOEmbedWithParams(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v1", "oembed", "", "")

	defer closeAPIServer()

	expectedResults := &OEmbed{}
	responseData := constructResponseAndResults("oembed_get.json", expectedResults)

	Convey("Given a URL and a OEmbedParams struct", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.OEmbed.Get("http://testing", OEmbedParams{Format: OEmbedFormatJSON})

			Convey("Then the oEmbed information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("format"), ShouldEqual, "json")
					So(receivedRequest.URL.Query().Get("url"), ShouldEqual, "http://testing")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, testResourcePath)
				})
			})
		})
	})
}
