package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

const testItemID = "16180339887"

func TestGetAllItems(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	expectedResults := ListItems{}
	responseData := constructResponseAndResults("items_get_all.json", &expectedResults)

	Convey("Given a board ID", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Items.GetAll(testBoardID)

			Convey("Then a slice of item information is returned", func() {
				So(err, ShouldBeNil)
				So(results.Data[0].ID, ShouldResemble, expectedResults.Data[0].ID)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/items", endpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestGetAllItemsWithSearchParams(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	expectedResults := ListItems{}
	responseData := constructResponseAndResults("items_get_all.json", &expectedResults)
	//roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a board ID", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Items.GetAll(testBoardID, ItemSearchParams{Type: ItemTypeFrame, Limit: "1"})

			Convey("Then a slice of item information is returned", func() {
				So(err, ShouldBeNil)
				So(results.Data[0].ID, ShouldResemble, expectedResults.Data[0].ID)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("type"), ShouldEqual, ItemTypeFrame)
					So(receivedRequest.URL.Query().Get("limit"), ShouldEqual, "1")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/items", endpointBoards, testBoardID))

					//Convey("And round-tripping the data does not result in any loss of data", func() {
					//	So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					//})
				})
			})
		})
	})
}

func TestGetItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	expectedResults := &Item{}
	responseData := constructResponseAndResults("items_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Get function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Items.Get(testBoardID, testItemID)

			Convey("Then a slice of item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/items/%s", endpointBoards, testBoardID, testItemID))
				})
			})
		})
	})
}

func TestUpdateItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	responseBody := &Item{}
	constructResponseAndResults("items_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a ItemUpdate struct", t, func() {
		Convey("When the Update function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := ItemUpdate{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Position.X = bodyData.Position.X
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.Items.Update(testBoardID, testItemID, ItemUpdate{Position: PositionSet{X: -1.5643}})

			Convey("Then the board member information is returned which includes the new role", func() {
				So(err, ShouldBeNil)
				So(results.Position.X, ShouldEqual, -1.5643)
				So(results.Position.Y, ShouldEqual, 100)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPatch)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/items/%s", endpointBoards, testBoardID, testItemID))
				})
			})
		})
	})

}

func TestDeleteItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.Items.Delete(testBoardID, testItemID)

			Convey("Then the item is deleted (no error is returned)", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodDelete)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/items/%s", endpointBoards, testBoardID, testItemID))
				})
			})
		})
	})
}
