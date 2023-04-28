package miro

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"time"
)

const (
	testBoardName     = "test-name"
	testBoardViewLink = "https://testing-miro.com/app/board"
	testBoardDesc     = "MIRO Gopher"
	testBoardID       = "3141592"
	testTeamID        = "662607015"
)

func boardResponse(boardID, description, teamID string, timeStamp time.Time) Board {
	return Board{
		ID:          boardID,
		Name:        testBoardName,
		ViewLink:    fmt.Sprintf("%s/%s", testBoardViewLink, boardID),
		Description: description,
		CreatedAt:   timeStamp,
		ModifiedAt:  timeStamp,
		Team: BasicEntityInfo{
			ID: teamID,
		},
	}
}

func getTimeNow() time.Time {
	timeStamp, _ := time.Parse("{2006-01-02 15:04:05.999999 -0700 MST}",
		time.Now().Format("{2006-01-02 15:04:05.999999 -0700 MST}"))
	return timeStamp
}

func TestCreateBoard(t *testing.T) {
	client, _, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, "", "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	testData := CreateBoard{
		Description: testBoardDesc,
		Name:        testBoardName,
		TeamID:      testTeamID,
	}
	expectedResults := boardResponse(testBoardID, testBoardDesc, testTeamID, timeStamp)

	Convey("Given a CreateBoard struct", t, func() {
		Convey("When the Boards Create function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				boardCreateData := CreateBoard{}
				json.NewDecoder(r.Body).Decode(&boardCreateData)

				json.NewEncoder(w).Encode(boardResponse(testBoardID, boardCreateData.Description, boardCreateData.TeamID, timeStamp))
				receivedRequest = r
			})

			results, err := client.Boards.Create(testData)

			Convey("Then the board is created and the board data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPost)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s", EndpointBoards))
				})
			})
		})
	})
}

func TestGetBoard(t *testing.T) {
	client, testURL, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, testBoardID, "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	expectedResults := boardResponse(testBoardID, testBoardDesc, testTeamID, timeStamp)

	Convey("Given a board ID", t, func() {
		Convey("When the Boards Get function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testURL, func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(boardResponse(testBoardID, testBoardDesc, testTeamID, timeStamp))
				receivedRequest = r
			})

			results, err := client.Boards.Get(testBoardID)

			Convey("Then the board data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s", EndpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestGetBoardWithError(t *testing.T) {
	client, testURL, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, "", "")
	defer closeAPIServer()

	tests := []struct {
		id          string
		responseErr ResponseError
		expectedErr string
	}{
		{
			id: "1",
			responseErr: ResponseError{
				Status:  http.StatusNotFound,
				Code:    "4.0101",
				Message: "Board not found",
				Type:    "error",
			},
			expectedErr: "Board not found",
		},
		{
			id: "2",
			responseErr: ResponseError{
				Status:  http.StatusBadRequest,
				Message: "Bad Request",
				Type:    "error",
			},
			expectedErr: "Bad Request",
		},
		{
			id: "3",
			responseErr: ResponseError{
				Status:  http.StatusTooManyRequests,
				Message: "Too many requests",
				Type:    "error",
			},
			expectedErr: "Too many requests",
		},
	}

	Convey("Given a board ID", t, func() {
		for _, test := range tests {
			Convey(fmt.Sprintf("When the Boards Get function is called and the expected return error is %s", test.expectedErr), func() {
				mux.HandleFunc(fmt.Sprintf("%s/%s", testURL, test.id), func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode(test.responseErr)
				})

				_, err := client.Boards.Get(test.id)

				Convey("Then the error is returned", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, test.expectedErr)
				})
			})
		}
	})
}

func TestListBoards(t *testing.T) {
	client, _, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, "", "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	expectedResults := &ListBoards{
		Total:  123,
		Size:   123,
		Offset: 1,
		Limit:  1,
		Data: []*Board{
			{
				CreatedAt:   timeStamp,
				CreatedBy:   BasicEntityInfo{ID: "2718282", Name: "Leonhard Euler", Type: "user"},
				Description: testBoardDesc,
				ID:          testBoardID,
				ModifiedAt:  timeStamp,
				Name:        testBoardName,
			},
		},
		Type: "board",
	}

	Convey("Given no arguments", t, func() {
		Convey("When the Boards GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(expectedResults)
				receivedRequest = r
			})

			results, err := client.Boards.GetAll()

			Convey("Then the list of boards is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s", EndpointBoards))
				})
			})
		})
	})
}

func TestListBoardsWithSearchParams(t *testing.T) {
	client, _, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, "", "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	expectedResults := &ListBoards{
		Total:  2,
		Size:   123,
		Offset: 1,
		Limit:  1,
		Data: []*Board{
			{
				CreatedAt:   timeStamp,
				CreatedBy:   BasicEntityInfo{ID: "1", Name: "Josef K", Type: "user"},
				Description: "test",
				ID:          "1",
				ModifiedAt:  timeStamp,
				Name:        "test",
				Owner:       BasicEntityInfo{ID: "30744573567", Name: "Franz Kafka", Type: "user"},
			},
			{
				CreatedAt:   timeStamp,
				CreatedBy:   BasicEntityInfo{ID: "2", Name: "Anna", Type: "user"},
				Description: "test",
				ID:          "2",
				ModifiedAt:  timeStamp,
				Name:        "test",
				Owner:       BasicEntityInfo{ID: "30744573567", Name: "Franz Kafka", Type: "user"},
			},
		},
		Type: "board",
	}

	Convey("Given a query string", t, func() {
		Convey("When the Boards GetAll function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(fmt.Sprintf("/v2/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
				reverseSlice(expectedResults.Data)
				json.NewEncoder(w).Encode(expectedResults)
				receivedRequest = r
			})

			results, err := client.Boards.GetAll(BoardSearchParams{
				Owner: "30744573567",
				Sort:  SortAlphabetically,
			})

			Convey("Then the list of boards is returned, sorted alphabetically", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, expectedResults)
				So(results.Data[0].CreatedBy.Name, ShouldEqual, "Anna")
				So(results.Data[1].CreatedBy.Name, ShouldEqual, "Josef K")

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodGet)
					So(receivedRequest.URL.Query().Get("sort"), ShouldEqual, "alphabetically")
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s", EndpointBoards))
				})
			})
		})
	})
}

func TestCopyBoard(t *testing.T) {
	client, testURL, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, "", "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	testData := CreateBoard{
		Description: testBoardDesc,
		Name:        testBoardName,
		TeamID:      testTeamID,
		Policy: Policy{
			SharingPolicy: SharingPolicy{
				Access:                            AccessPrivate,
				InviteToAccountAndBoardLinkAccess: InviteAccessEditor,
				OrganizationAccess:                AccessEdit,
				TeamAccess:                        AccessEdit,
			},
			PermissionsPolicy: PermissionsPolicy{
				SharingAccess:                 SharingAccessOwnersAndCoOwners,
				CopyAccess:                    CopyAccessTeamEditors,
				CollaborationToolsStartAccess: CollabAccessBoardOwnersAndCoOwners,
			},
		},
	}
	expectedResults := boardResponse(testBoardID, testBoardDesc, testTeamID, timeStamp)

	var receivedRequest *http.Request
	mux.HandleFunc(testURL, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("copy_from") != "" {
			w.WriteHeader(http.StatusCreated)
			boardCreateData := CreateBoard{}
			json.NewDecoder(r.Body).Decode(&boardCreateData)
			json.NewEncoder(w).Encode(boardResponse(testBoardID, boardCreateData.Description, boardCreateData.TeamID, timeStamp))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseError{
				Status:  http.StatusBadRequest,
				Message: "invalid request",
			})
		}
		receivedRequest = r
	})

	Convey("Given a CreateBoard struct", t, func() {
		Convey(fmt.Sprintf("When the Boards Copy function is called with valid data"), func() {
			results, err := client.Boards.Copy(testData, testBoardID)

			Convey("Then the board is created and the board data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResults)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPut)
					So(receivedRequest.URL.Query().Get("copy_from"), ShouldEqual, testBoardID)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s", EndpointBoards))
				})
			})
		})
		Convey(fmt.Sprintf("When the Boards Copy function is called with invalid data"), func() {
			results, err := client.Boards.Copy(testData, "")

			Convey("Then the board is not created and an error is returned", func() {
				So(err, ShouldBeError)
				So(results, ShouldResemble, &Board{})
			})
		})
	})
}

func TestUpdateBoard(t *testing.T) {
	client, testURL, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, testBoardID, "")
	defer closeAPIServer()

	timeStamp := getTimeNow()

	testData := CreateBoard{
		Description: "A new description",
		Name:        testBoardName,
		TeamID:      testTeamID,
	}
	expectedResult := boardResponse(testBoardID, "A new description", testTeamID, timeStamp)

	Convey("Given a CreateBoard struct", t, func() {
		Convey("When the Boards Update function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testURL, func(w http.ResponseWriter, r *http.Request) {
				boardCreateData := CreateBoard{}
				json.NewDecoder(r.Body).Decode(&boardCreateData)

				json.NewEncoder(w).Encode(boardResponse(testBoardID, boardCreateData.Description, boardCreateData.TeamID, timeStamp))
				receivedRequest = r
			})

			results, err := client.Boards.Update(testData, testBoardID)

			Convey("Then the board is updated and the board data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, &expectedResult)
				So(results.Description, ShouldEqual, expectedResult.Description)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodPatch)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s", EndpointBoards, testBoardID))
				})
			})
		})
	})
}

func TestDeleteBoard(t *testing.T) {
	client, testURL, mux, closeAPIServer := mockMIROAPI("v2", EndpointBoards, testBoardID, "")
	defer closeAPIServer()

	Convey("Given a board ID", t, func() {
		Convey("When the Boards Delete function is called", func() {
			var receivedRequest *http.Request
			mux.HandleFunc(testURL, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				receivedRequest = r
			})

			err := client.Boards.Delete(testBoardID)

			Convey("Then the board is deleted (no error is returned)", func() {
				So(err, ShouldBeNil)

				Convey("And the request contains the expected headers and parameters", func() {
					So(receivedRequest, ShouldNotBeNil)
					So(receivedRequest.Method, ShouldEqual, http.MethodDelete)
					So(receivedRequest.Header.Get("Authorization"), ShouldEqual, fmt.Sprintf("Bearer %s", testToken))
					So(receivedRequest.URL.Path, ShouldEqual, fmt.Sprintf("/v2/%s/%s", EndpointBoards, testBoardID))
				})
			})
		})
	})
}

func reverseSlice(s []*Board) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
