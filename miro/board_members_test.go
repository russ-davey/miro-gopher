package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

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
