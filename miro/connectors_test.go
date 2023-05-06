package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestCreateConnector(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	expectedResults := &Connector{}
	responseData := constructResponseAndResults("connectors_get.json", &expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Create method is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Connectors.Create(testBoardID,
				SetConnector{})

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

func TestGetConnector(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	expectedResults := Connector{}
	responseData := constructResponseAndResults("connectors_get.json", &expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Connectors.Get(testBoardID, testItemID)

			Convey("Then the item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/connectors/%s", endpointBoards, testBoardID, testItemID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestGetAllConnectors(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	expectedResults := &ListConnectors{}
	responseData := constructResponseAndResults("connectors_get_all.json", expectedResults)

	Convey("Given no arguments", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Connectors.GetAll(testBoardID)

			Convey("Then a list of item information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/connectors", endpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestGetAllConnectorsWithSearchParams(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	expectedResults := &ListConnectors{}
	responseData := constructResponseAndResults("connectors_get_all.json", expectedResults)

	Convey("Given a search param to limit the number of results returned", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.Connectors.GetAll(testBoardID, ConnectorSearchParams{Limit: "1"})

			Convey("Then a list of item information is returned consisting of just one item", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("limit"), ShouldEqual, "1")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/connectors", endpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestUpdateConnector(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	responseBody := Connector{}
	constructResponseAndResults("connectors_get.json", &responseBody)

	Convey("Given a board ID, an item ID and a SetConnector struct", t, func() {
		Convey("When Update is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := SetConnector{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Shape = bodyData.Shape
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.Connectors.Update(testBoardID, testItemID, SetConnector{Shape: ConnectorShapeElbowed})

			Convey("Then the item information is returned which includes the new shape", func() {
				So(err, ShouldBeNil)
				So(results.Shape, ShouldEqual, "elbowed")

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPatch)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/connectors/%s", endpointBoards, testBoardID, testItemID))
				})
			})
		})
	})
}

func TestDeleteConnector(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "connectors")
	defer closeAPIServer()

	Convey("Given a board ID and an item ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testItemID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.Connectors.Delete(testBoardID, testItemID)

			Convey("Then the item is deleted (no error is returned)", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodDelete)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/connectors/%s", endpointBoards, testBoardID, testItemID))
				})
			})
		})
	})
}
