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

			results, err := client.DocumentItems.Create(testBoardID, DocumentItemSet{
				Data: ItemDataSet{Title: "A test", URL: "http://testing"},
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

func TestUploadDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	responseBody := &DocumentItem{}
	constructResponseAndResults("document_item_get.json", &responseBody)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Upload method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				r.ParseMultipartForm(32 << 20)

				file, _, _ := r.FormFile("data")
				defer file.Close()

				bodyData := UploadFileItem{}
				json.NewDecoder(file).Decode(&bodyData)

				responseBody.Data.Title = bodyData.Title
				jsonData, _ := json.Marshal(responseBody)

				w.WriteHeader(http.StatusCreated)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.DocumentItems.Upload(testBoardID, "./test_data/document_item_get.json", UploadFileItem{Title: "A test upload"})

			Convey("Then the item is uploaded", func() {
				So(err, ShouldBeNil)
				So(results.Data.Title, ShouldEqual, "A test upload")

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

	Convey("Given a board ID, an item ID and a DocumentItemSet struct", t, func() {
		Convey("When the Update function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := DocumentItemSet{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Data.Title = bodyData.Data.Title
				responseBody.Data.DocumentURL = bodyData.Data.URL
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.DocumentItems.Update(testBoardID, testItemID, DocumentItemSet{
				Data: ItemDataSet{Title: "A test", URL: "http://testing"}})

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

func TestUploadUpdateDocumentItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "documents")
	defer closeAPIServer()

	responseBody := &DocumentItem{}
	constructResponseAndResults("document_item_get.json", &responseBody)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the UpdateFromFile method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				r.ParseMultipartForm(32 << 20)

				file, _, _ := r.FormFile("data")
				defer file.Close()

				bodyData := UploadFileItem{}
				json.NewDecoder(file).Decode(&bodyData)

				responseBody.Data.Title = bodyData.Title
				jsonData, _ := json.Marshal(responseBody)

				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.DocumentItems.UpdateFromFile(testBoardID, testItemID, "./test_data/document_item_get.json", UploadFileItem{Title: "A test upload update"})

			Convey("Then the item is uploaded", func() {
				So(err, ShouldBeNil)
				So(results.Data.Title, ShouldEqual, "A test upload update")

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
