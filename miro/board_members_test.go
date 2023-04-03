package miro

//func TestShareBoard(t *testing.T) {
//	client, mux, _, closeAPIServer := mockMIROAPI()
//	defer closeAPIServer()
//
//	tests := map[string]struct {
//		id       string
//		emails   []string
//		expected *ListBoardsResponse
//	}{
//		ListBoardsResponse{},
//		"ok": {"1", []string{"keke@miro.com", "miro@keke.com"}, getBoardList()},
//	}
//
//	Convey("Given a board ID", t, func() {
//		for _, test := range tests {
//			Convey("When the ListBoards Share function is called", func() {
//				mux.HandleFunc(fmt.Sprintf("/boards/%s/share", test.id), func(w http.ResponseWriter, r *http.Request) {
//					json.NewEncoder(w).Encode(boardResponse(test.id, timeStamp))
//				})
//
//				results, err := client.Boards.Get(test.id)
//				if err != nil {
//					t.Fatalf("Failed: %v", err)
//				}
//
//				Convey("Then the board data is returned", func() {
//					So(results, ShouldResemble, test.expected)
//				})
//			})
//		}
//	})
//}
