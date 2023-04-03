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
	testBoardViewLink = "https://test-test.com"
	testBoardDesc     = ""
)

func boardResponse(id string, timeStamp time.Time) Board {
	return Board{
		ID:          id,
		Name:        testBoardName,
		ViewLink:    testBoardViewLink,
		Description: testBoardDesc,
		CreatedAt:   timeStamp,
		ModifiedAt:  timeStamp,
	}
}

func getBoard(id string, timeStamp time.Time) *Board {
	return &Board{
		ID:          id,
		ViewLink:    testBoardViewLink,
		Name:        testBoardName,
		Description: testBoardDesc,
		ModifiedAt:  timeStamp,
		CreatedAt:   timeStamp,
	}
}

func getTimeNow() time.Time {
	timeStamp, _ := time.Parse("{2006-01-02 15:04:05.999999 -0700 MST}",
		time.Now().Format("{2006-01-02 15:04:05.999999 -0700 MST}"))
	return timeStamp
}

func TestCreateBoard(t *testing.T) {
	client, mux, _, closeAPIServer := mockMIROAPI()
	defer closeAPIServer()

	timeStamp := getTimeNow()

	tests := []struct {
		id       string
		body     CreateBoard
		expected *Board
	}{
		{
			id: "1",
			body: CreateBoard{
				Description: testBoardDesc,
				Name:        testBoardName,
				TeamID:      "3141592",
			},
			expected: getBoard("1", timeStamp),
		},
	}

	Convey("Given a CreateBoard struct", t, func() {
		for _, test := range tests {
			Convey("When the Boards Create function is called", func() {
				mux.HandleFunc(fmt.Sprintf("/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
					if r.Method == http.MethodPost {
						w.WriteHeader(http.StatusCreated)
						json.NewEncoder(w).Encode(boardResponse(test.id, timeStamp))
					}
				})

				results, err := client.Boards.Create(test.body)

				Convey("Then the board is created and the board data is returned", func() {
					So(err, ShouldBeNil)
					So(results, ShouldResemble, test.expected)
				})
			})
		}
	})
}

func TestGetBoard(t *testing.T) {
	client, mux, _, closeAPIServer := mockMIROAPI()
	defer closeAPIServer()

	timeStamp := getTimeNow()

	tests := []struct {
		id       string
		expected *Board
	}{
		{
			id:       "1",
			expected: getBoard("1", timeStamp),
		},
	}

	Convey("Given a board ID", t, func() {
		for _, test := range tests {
			Convey("When the Boards Get function is called", func() {
				mux.HandleFunc(fmt.Sprintf("/%s/%s", EndpointBoards, test.id), func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(boardResponse(test.id, timeStamp))
				})

				results, err := client.Boards.Get(test.id)

				Convey("Then the board data is returned", func() {
					So(err, ShouldBeNil)
					So(results, ShouldResemble, test.expected)
				})
			})
		}
	})
}

func TestGetBoardWithError(t *testing.T) {
	client, mux, _, closeAPIServer := mockMIROAPI()
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
				mux.HandleFunc(fmt.Sprintf("/%s/%s", EndpointBoards, test.id), func(w http.ResponseWriter, r *http.Request) {
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
	client, mux, _, closeAPIServer := mockMIROAPI()
	defer closeAPIServer()

	timeStamp := getTimeNow()

	tests := []struct {
		query    string
		expected *ListBoards
	}{
		{
			query: "test",
			expected: &ListBoards{
				Data: []*BoardData{
					{
						CreatedAt: timeStamp,
						CreatedBy: BasicUserInfo{
							ID:   "123",
							Name: "John Smith",
							Type: "user",
						},
						Description: "test",
						ID:          "1",
						ModifiedAt:  timeStamp,
						Name:        "test",
					},
				},
				Total:  123,
				Size:   123,
				Offset: 1,
				Limit:  1,
			},
		},
	}

	Convey("Given a query string", t, func() {
		for _, test := range tests {
			Convey("When the Boards GetAll function is called", func() {
				mux.HandleFunc(fmt.Sprintf("/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(test.expected)
				})

				results, err := client.Boards.GetAll()

				Convey("Then the list of boards is returned", func() {
					So(err, ShouldBeNil)
					So(results, ShouldResemble, test.expected)
				})
			})
		}
	})
}

func TestListBoardsWithQueryParams(t *testing.T) {
	client, mux, _, closeAPIServer := mockMIROAPI()
	defer closeAPIServer()

	timeStamp := getTimeNow()

	tests := []struct {
		query    string
		expected *ListBoards
	}{
		{
			query: "test",
			expected: &ListBoards{
				Data: []*BoardData{
					{
						CreatedAt: timeStamp,
						CreatedBy: BasicUserInfo{
							ID:   "123",
							Name: "Franz Kafka",
							Type: "user",
						},
						Description: "test",
						ID:          "1",
						ModifiedAt:  timeStamp,
						Name:        "test",
						Owner: BasicUserInfo{
							ID:   "30744573567",
							Name: "Franz Kafka",
							Type: "user",
						},
					},
					{
						CreatedAt: timeStamp,
						CreatedBy: BasicUserInfo{
							ID:   "123",
							Name: "Anna",
							Type: "user",
						},
						Description: "test",
						ID:          "1",
						ModifiedAt:  timeStamp,
						Name:        "test",
						Owner: BasicUserInfo{
							ID:   "30744573567",
							Name: "Anna",
							Type: "user",
						},
					},
				},
				Total:  123,
				Size:   123,
				Offset: 1,
				Limit:  1,
			},
		},
	}

	Convey("Given a query string", t, func() {
		for _, test := range tests {
			Convey("When the Boards GetAll function is called", func() {
				mux.HandleFunc(fmt.Sprintf("/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Query().Get("sort") == "alphabetically" {
						reverseSlice(test.expected.Data)
					}
					json.NewEncoder(w).Encode(test.expected)
				})

				results, err := client.Boards.GetAll(BoardQueryParams{
					Owner: "30744573567",
					Sort:  SortAlphabetically,
				})

				Convey("Then the list of boards is returned, sorted alphabetically", func() {
					So(err, ShouldBeNil)
					So(results, ShouldResemble, test.expected)
					So(results.Data[0].Owner.Name, ShouldEqual, "Anna")
					So(results.Data[1].Owner.Name, ShouldEqual, "Franz Kafka")
				})
			})
		}
	})
}

func TestCopyBoard(t *testing.T) {
	testID := "1"
	timeStamp := getTimeNow()
	testBody := CreateBoard{
		Description: testBoardDesc,
		Name:        testBoardName,
		TeamID:      "3141592",
		Policy: Policy{
			SharingPolicy: SharingPolicy{
				Access:                            AccessPrivate,
				InviteToAccountAndBoardLinkAccess: InviteAccessEditor,
				OrganizationAccess:                AccessEdit,
				TeamAccess:                        AccessEdit,
			},
			PermissionsPolicy: PermissionsPolicy{
				SharingAccess:                 SharingAccessTeamMemberWithEditingRights,
				CopyAccess:                    CopyAccessTeamEditors,
				CollaborationToolsStartAccess: AccessBoardOwnersAndCoOwners,
			},
		},
	}
	TestQueryParams := "123456"

	client, mux, _, closeAPIServer := mockMIROAPI()
	defer closeAPIServer()

	mux.HandleFunc(fmt.Sprintf("/%s", EndpointBoards), func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			if r.URL.Query().Get("copy_from") != "" {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(boardResponse(testID, timeStamp))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ResponseError{
					Status:  http.StatusBadRequest,
					Message: "invalid request",
				})
			}
		}
	})

	Convey("Given a CreateBoard struct", t, func() {
		Convey(fmt.Sprintf("When the Boards Copy function is called with valid data"), func() {
			results, err := client.Boards.CopyBoard(testBody, TestQueryParams)

			Convey("Then the board is created and the board data is returned", func() {
				So(err, ShouldBeNil)
				So(results, ShouldResemble, getBoard("1", timeStamp))
			})
		})
		Convey(fmt.Sprintf("When the Boards Copy function is called with invalid data"), func() {
			results, err := client.Boards.CopyBoard(testBody, "")

			Convey("Then the board is not created and an error is returned", func() {
				So(err, ShouldBeError)
				So(results, ShouldResemble, &Board{})
			})
		})
	})
}

func reverseSlice(s []*BoardData) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

//func TestGetBoardForReal(t *testing.T) {
//	client := NewClient(os.Getenv("MIRO_TOKEN"))
//
//	//board, err := client.Boards.Get("uXjVMXqSjzc=")
//	board, err := client.Boards.CopyBoard(CreateBoard{})
//	if err != nil {
//		fmt.Printf("Error: %v\n", err)
//	}
//
//	jsonData, _ := json.Marshal(board)
//	fmt.Println(string(jsonData))
//}
