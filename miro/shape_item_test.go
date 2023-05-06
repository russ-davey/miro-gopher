package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestCreateShapeItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, "uXjVMNoCEUs=", "shapes")
	defer closeAPIServer()

	expectedResults := &ShapeItem{}
	responseData := constructResponseAndResults("shape_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Create method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.ShapeItems.Create("uXjVMNoCEUs=", SetShapeItem{
				Data: ShapeItemData{
					Shape:   ShapeTriangle,
					Content: "Bill Cipher",
				},
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

func TestGetShapeItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, "uXjVMNoCEUs=", "shapes")
	defer closeAPIServer()

	expectedResults := &ShapeItem{}
	responseData := constructResponseAndResults("shape_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.ShapeItems.Get("uXjVMNoCEUs=", testItemID)

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

func TestUpdateShapeItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, "uXjVMNoCEUs=", "shapes")
	defer closeAPIServer()

	responseBody := &ShapeItem{}
	constructResponseAndResults("shape_item_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a SetShapeItem struct", t, func() {
		Convey("When Update is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := SetShapeItem{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Data = bodyData.Data
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.ShapeItems.Update("uXjVMNoCEUs=", testItemID, SetShapeItem{Data: ShapeItemData{
				Shape: "triangle", Content: "Bill Cipher"}})

			Convey("Then the item information is returned which includes the new data", func() {
				So(err, ShouldBeNil)
				So(results.Data.Shape, ShouldEqual, "triangle")
				So(results.Data.Content, ShouldEqual, "Bill Cipher")

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

func TestDeleteShapeItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, "uXjVMNoCEUs=", "shapes")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.ShapeItems.Delete("uXjVMNoCEUs=", testItemID)

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
