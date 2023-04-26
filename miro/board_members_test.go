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
	client, mux, closeAPIServer := mockMIROAPI("v2")
	defer closeAPIServer()

	Convey("Given a Board ID", t, func() {
		Convey("When the ShareBoard function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(BoardInvitationResponse{
					Successful: 3074457350804038700,
				})
				receivedRequest = r
			})

			results, err := client.BoardMembers.ShareBoard(ShareBoardInvitation{
				Emails:  []string{"alpha@milkyway.com", "beta@milkyway.com", "gramma@milkyway.com", "delta@milkyway.com"},
				Role:    RoleEditor,
				Message: "Join me on my awesome board",
			}, testBoardID)

			Convey("Then a success message response is returned", func() {
				So(err, ShouldBeNil)
				So(results.Successful, ShouldEqual, 3074457350804038700)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPost)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestGetBoardMember(t *testing.T) {
	client, mux, closeAPIServer := mockMIROAPI("v2")
	defer closeAPIServer()

	expectedResults := BoardMember{}
	responseData := constructResponseAndResults("board_members_get.json", &expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given a board ID and a board member ID", t, func() {
		Convey("When the Get function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s/%s/members/%s", EndpointBoards, testBoardID, testBoardMemberID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.Get(testBoardID, testBoardMemberID)

			Convey("Then a slice of board member information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members/%s", EndpointBoards, testBoardID, testBoardMemberID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestGetAllBoardMembers(t *testing.T) {
	client, mux, closeAPIServer := mockMIROAPI("v2")
	defer closeAPIServer()

	expectedResults := &ListBoardMembersResponse{}
	responseData := constructResponseAndResults("board_members_get_all.json", expectedResults)
	roundTrip, _ := json.Marshal(expectedResults)

	Convey("Given no arguments", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.GetAll(testBoardID)

			Convey("Then a slice of board member information is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID))

					Convey("And round-tripping the data does not result in any loss of data", func() {
						So(compareJSON(responseData, roundTrip), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestGetAllBoardMembersWithSearchParams(t *testing.T) {
	client, mux, closeAPIServer := mockMIROAPI("v2")
	defer closeAPIServer()

	expectedResults := &ListBoardMembersResponse{}
	responseData := constructResponseAndResults("board_members_get_all.json", expectedResults)

	Convey("Given a search param to limit the number of results returned", t, func() {
		Convey("When the GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID), func(w http.ResponseWriter, r *http.Request) {
				w.Write(responseData)
				receivedRequest = r
			})

			results, err := client.BoardMembers.GetAll(testBoardID, BoardMemberSearchParams{limit: "1"})

			Convey("Then a slice of board member information is returned consisting of just one member", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("limit"), ShouldEqual, "1")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s/members", EndpointBoards, testBoardID))
				})
			})
		})
	})
}
