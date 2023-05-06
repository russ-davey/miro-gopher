package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

const testBoardMemberID = "2718282"

func TestShareBoard(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	Convey("Given a Board ID", t, func() {
		Convey("When ShareBoard is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(BoardInvitationResponse{
					Successful: "3074457350804038700",
				})
				receivedRequest = r
			})

			results, err := client.BoardMembers.ShareBoard(testBoardID, ShareBoardInvitation{
				Emails:  []string{"alpha@quadrant.com", "beta@quadrant.com", "gramma@quadrant.com", "delta@quadrant.com"},
				Role:    RoleEditor,
				Message: "Join me on my awesome board, make it so",
			})

			Convey("Then a success message response is returned", func() {
				So(err, ShouldBeNil)
				So(results.Successful, ShouldEqual, "3074457350804038700")

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPost)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", endpointBoards, testBoardID))
				})
			})
		})

		Convey("When ShareBoard is called without a board ID", func() {
			_, err := client.BoardMembers.ShareBoard("", ShareBoardInvitation{})

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetBoardMember(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	expectedResults := BoardMember{}
	responseData := constructResponseAndResults("board_members_get.json", &expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a board ID and a board member ID", t, func() {
		Convey("When Get is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testBoardMemberID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.Get(testBoardID, testBoardMemberID)

			Convey("Then the board member information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members/%s", endpointBoards, testBoardID, testBoardMemberID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})

		Convey("When Get is called without a board ID and an item ID", func() {
			_, err := client.BoardMembers.Get("", "")

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetAllBoardMembers(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	expectedResults := &ListBoardMembers{}
	responseData := constructResponseAndResults("board_members_get_all.json", expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given no arguments", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.GetAll(testBoardID)

			Convey("Then a list of board member information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", endpointBoards, testBoardID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})

		Convey("When GetAll is called without a board ID", func() {
			_, err := client.BoardMembers.GetAll("")

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestGetAllBoardMembersWithSearchParams(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	expectedResults := &ListBoardMembers{}
	responseData := constructResponseAndResults("board_members_get_all.json", expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a search param to limit the number of results returned", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testResourcePath, func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.GetAll(testBoardID, BoardMemberSearchParams{Limit: "1"})

			Convey("Then a list of board member information is returned consisting of just one member", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("limit"), ShouldEqual, "1")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", endpointBoards, testBoardID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})

		Convey("When GetAll is called without a board ID", func() {
			_, err := client.BoardMembers.GetAll("", BoardMemberSearchParams{Limit: "1"})

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestUpdateBoardMember(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	responseBody := BoardMember{}
	constructResponseAndResults("board_members_get.json", &responseBody)

	Convey("Given a board ID, a board member ID and a new role", t, func() {
		Convey("When Update is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testBoardMemberID), func(w http.ResponseWriter, r *http.Request) {
				// decode body
				bodyData := RoleUpdate{}
				json.NewDecoder(r.Body).Decode(&bodyData)
				// update test data
				responseBody.Role = bodyData.Role
				// marshal test data
				jsonData, _ := json.Marshal(responseBody)
				w.Write(jsonData)

				receivedRequest = r
			})

			results, err := client.BoardMembers.Update(testBoardID, testBoardMemberID, RoleEditor)

			Convey("Then the board member information is returned which includes the new role", func() {
				So(err, ShouldBeNil)
				So(results.Role, ShouldEqual, RoleEditor)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPatch)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members/%s", endpointBoards, testBoardID, testBoardMemberID))
				})
			})
		})

		Convey("When Update is called without a board ID and an item ID", func() {
			_, err := client.BoardMembers.Update("", "", RoleEditor)

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestDeleteBoardMember(t *testing.T) {
	client, testResourcePath, mux, closeAPIServer := mockMIROAPI("v2", endpointBoards, testBoardID, "members")
	defer closeAPIServer()

	Convey("Given a board ID and a board member ID", t, func() {
		Convey("When the Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("%s/%s", testResourcePath, testBoardMemberID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.BoardMembers.Delete(testBoardID, testBoardMemberID)

			Convey("Then the board member is deleted (no error is returned)", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodDelete)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members/%s", endpointBoards, testBoardID, testBoardMemberID))
				})
			})
		})

		Convey("When Delete is called without a board ID and an item ID", func() {
			err := client.BoardMembers.Delete("", "")

			Convey("Then an error is returned", func() {
				So(err, ShouldBeError)
			})
		})
	})
}
