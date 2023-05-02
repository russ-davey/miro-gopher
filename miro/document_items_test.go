package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestCreateDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	expectedResults := &DocumentItem{}
	responseData := constructResponseAndResults("document_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Create method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.DocumentItems.Create(testBoardID, SetDocumentItem{
				Data: DocumentItemData{Title: "A test", DocumentURL: "http://testing"},
			})

			Convey("Then the item is created", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPost)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, testResourcePath)
				})
			})
		})
	})
}

func TestGetDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	expectedResults := &DocumentItem{}
	responseData := constructResponseAndResults("document_item_get.json", &expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Get function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.DocumentItems.Get(testBoardID, testItemID)

			Convey("Then the item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("%s/%s", testResourcePath, testItemID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestUpdateDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	responseBody := &DocumentItem{}
	constructResponseAndResults("document_item_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a SetDocumentItem struct", t, func() {
		Convey("When the Update function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := DocumentItem{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Data = bodyData.Data
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.DocumentItems.Update(testBoardID, testItemID, SetDocumentItem{
				Data: DocumentItemData{Title: "A test", DocumentURL: "http://testing"}})

			Convey("Then the item information is returned which includes the new role", func() {
				So(err, ShouldBeNil)
				So(results.Data.Title, ShouldEqual, "A test")
				So(results.Data.DocumentURL, ShouldEqual, "http://testing")

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPatch)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("%s/%s", testResourcePath, testItemID))
				})
			})
		})
	})

}

func TestDeleteDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.DocumentItems.Delete(testBoardID, testItemID)

			Convey("Then the item is deleted (no error is returned)", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodDelete)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("%s/%s", testResourcePath, testItemID))
				})
			})
		})
	})
}
