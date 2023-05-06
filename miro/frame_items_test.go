package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestCreateFrame(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "frames")
	defer closeAPIServer()

	expectedResults := &FrameItem{}
	responseData := constructResponseAndResults("app_card_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Create method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Frames.Create(testBoardID,
				SetFrameItem{})

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

func TestGetFrame(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "frames")
	defer closeAPIServer()

	expectedResults := FrameItem{}
	responseData := constructResponseAndResults("frame_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Frames.Get(testBoardID, testItemID)

			Convey("Then the item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

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

func TestGetFrameItems(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	expectedResults := &ListItems{}
	responseData := constructResponseAndResults("frame_item_get_all.json", expectedResults)

	Convey("Given no arguments", t, func() {
		Convey("When the GetItems function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Frames.GetItems(testBoardID, testItemID)

			Convey("Then a list of item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("parent_item_id"), ShouldEqual, testItemID)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, testResourcePath)
				})
			})
		})
	})
}

func TestGetAllFramesWithSearchParams(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "items")
	defer closeAPIServer()

	expectedResults := &ListItems{}
	responseData := constructResponseAndResults("frame_item_get_all.json", expectedResults)

	Convey("Given a search param to limit the number of results returned", t, func() {
		Convey("When the GetItems function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Frames.GetItems(testBoardID, testItemID, ItemSearchParams{Limit: "1"})

			Convey("Then a list of item information is returned consisting of just one item", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("limit"), ShouldEqual, "1")
					So(receivedRequest.URL.Query().Get("parent_item_id"), ShouldEqual, testItemID)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, testResourcePath)
				})
			})
		})
	})
}

func TestUpdateFrameItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "frames")
	defer closeAPIServer()

	responseBody := &FrameItem{}
	constructResponseAndResults("frame_item_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a FrameItemUpdate struct", t, func() {
		Convey("When Update is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := SetFrameItem{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Position.X = bodyData.Position.X
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.Frames.Update(testBoardID, testItemID, SetFrameItem{
				Position: PositionSet{X: -2.7315},
			})

			Convey("Then the item information is returned which includes the new role", func() {
				So(err, ShouldBeNil)
				So(results.Position.X, ShouldEqual, -2.7315)

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

func TestDeleteFrame(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "frames")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.Frames.Delete(testBoardID, testItemID)

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
