package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestCreateAppCardItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "app_cards")
	defer closeAPIServer()

	expectedResults := &AppCardItem{}
	responseData := constructResponseAndResults("app_card_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Create is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.AppCardItems.Create(testBoardID,
				AppCardItemSet{
					Data: AppCardItemData{
						Fields: []Fields{
							{
								IconShape: "round",
								FillColor: "#2fa9e3",
								IconUrl:   "https://cdn-icons-png.flaticon.com/512/5695/5695864.png",
								TextColor: "#1a1a1a",
								Tooltip:   "Tooltip",
								Value:     "Status: In Progress",
							},
						},
						Status: StatusConnected,
					},
					Style: Style{
						FillColor: "#2d9bf0",
					},
					Position: PositionSet{
						Origin: Center,
						X:      2.5,
						Y:      2.5,
					},
					Geometry: Geometry{
						Height:   60,
						Rotation: 1.25,
						Width:    320.4,
					},
					Parent: ParentSet{
						ID: "123214124",
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

		Convey("When Create is called without a board ID", func() {
			_, err := client.AppCardItems.Create("", AppCardItemSet{})

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetAppCardItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "app_cards")
	defer closeAPIServer()

	expectedResults := &AppCardItem{}
	responseData := constructResponseAndResults("app_card_item_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.AppCardItems.Get(testBoardID, testItemID)

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

		Convey("When Get is called without a board ID or item ID", func() {
			_, err := client.AppCardItems.Get("", "")

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestUpdateAppCardItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "app_cards")
	defer closeAPIServer()

	responseBody := &AppCardItem{}
	constructResponseAndResults("app_card_item_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a AppCardItemUpdate struct", t, func() {
		Convey("When Update is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := AppCardItemSet{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Position.X = bodyData.Position.X
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.AppCardItems.Update(testBoardID, testItemID, AppCardItemSet{
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

		Convey("When Update is called without a board ID or item ID", func() {
			_, err := client.AppCardItems.Update("", "", AppCardItemSet{})

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})

}

func TestDeleteAppCardItem(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "app_cards")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Delete is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.AppCardItems.Delete(testBoardID, testItemID)

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

		Convey("When Delete is called without a board ID and an item ID", func() {
			err := client.AppCardItems.Delete("", "")

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}
