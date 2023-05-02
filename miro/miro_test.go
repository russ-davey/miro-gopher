package miro

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"
	"time"
)

const (
	testToken = "test-token"
)

func mockMIROAPI(apiVersion, resource, pathParam, subResource string) (*Client, string, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	handler := http.NewServeMux()

	handler.Handle(fmt.Sprintf("/%s/", apiVersion), mux)

	server := httptest.NewServer(handler)

	client := NewClient(testToken)
	client.BaseURL = server.URL

	var resourcePath string
	if pathParam == "" {
		resourcePath = fmt.Sprintf("/%s/%s", apiVersion, resource)
	} else if subResource == "" {
		resourcePath = fmt.Sprintf("/%s/%s/%s", apiVersion, resource, pathParam)
	} else {
		resourcePath = fmt.Sprintf("/%s/%s/%s/%s", apiVersion, resource, pathParam, subResource)
	}

	return client, resourcePath, mux, server.Close
}

func headerExists(req *http.Request, header string) bool {
	if req.Header.Get(header) != "" {
		return true
	}
	return false
}

func constructResponseAndResults(testData string, expectedResults interface{}) []byte {
	responseData, err := os.ReadFile(fmt.Sprintf("./test_data/%s", testData))
	if err != nil {
		log.Fatalf("error reading test data: %v\n", err)
	}
	if err := json.Unmarshal(responseData, expectedResults); err != nil {
		log.Fatalf("error decoding test data: %v\n", err)
	}

	return responseData
}

// compareJSON validate structs by round-tripping the test data and comparing the original to the data unmarshalled/marshalled
func compareJSON(json1, json2 []byte) bool {
	sortedJSON1, sortedJSON2 := sortJSON(json1, json2)

	// Compare
	return bytes.Equal(sortedJSON1, sortedJSON2)
}

func sortJSON(json1, json2 []byte) ([]byte, []byte) {
	// Parse the JSON bytes into maps
	var map1 map[string]interface{}
	var map2 map[string]interface{}
	json.Unmarshal(json1, &map1)
	json.Unmarshal(json2, &map2)

	// Sort the maps by keys
	sortedMap1 := make(map[string]interface{})
	sortedMap2 := make(map[string]interface{})
	var keys1 []string
	var keys2 []string
	for k, v := range map1 {
		if k == "modifiedAt" || k == "createdAt" {
			if strValue, ok := v.(string); ok {
				t, _ := time.Parse(time.RFC3339, strValue)
				v = t.String()
			}
		}
		keys1 = append(keys1, k)
		sortedMap1[k] = v
	}
	for k, v := range map2 {
		if k == "modifiedAt" || k == "createdAt" {
			if strValue, ok := v.(string); ok {
				t, _ := time.Parse(time.RFC3339, strValue)
				v = t.String()
			}
		}
		keys2 = append(keys2, k)
		sortedMap2[k] = v
	}
	sort.Strings(keys1)
	sort.Strings(keys2)

	// Convert the sorted maps to JSON strings
	sortedJSON1, _ := json.Marshal(sortedMap1)
	sortedJSON2, _ := json.Marshal(sortedMap2)

	return sortedJSON1, sortedJSON2
}

//func TestStructAgainstRealData(t *testing.T) {
//	client := NewClient(os.Getenv("MIRO_TOKEN"))
//
//	//version := "v2"
//	//resource := "boards"
//	//subResource := "shapes"
//	boardID := "="
//	//itemID := ""
//
//	response, err := client.ShapeItems.Update(boardID, "", SetShapeItem{
//		Data: ShapeItemData{Shape: ShapeCloud, Content: "A fluffy cloud"},
//		Style: Style{
//			TextAlign: TextAlignCenter,
//			FontSize:  "10",
//		},
//		Position: PositionSet{Origin: Center, X: -100, Y: -100},
//		Geometry: Geometry{
//			Height:   1000,
//			Rotation: 90,
//			Width:    1000,
//		},
//	})
//	//response, err := client.ShapeItems.Get(boardID, itemID)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	jsonData, _ := json.Marshal(response)
//	fmt.Println(string(jsonData))
//
//	//rawResponse := make(map[string]interface{})
//	////client.Get(fmt.Sprintf("https://api.miro.com/%s/%s", version, endpoint), &rawResponse)
//	//client.Get(fmt.Sprintf("https://api.miro.com/%s/%s/%s/%s/%s", version, resource, boardID, subResource, itemID), &rawResponse)
//	//jsonDataNative, _ := json.Marshal(rawResponse)
//	//
//	//processed, native := sortJSON(jsonData, jsonDataNative)
//	//fmt.Printf("== Processed ==: %s\n", processed)
//	//fmt.Println("===============================")
//	//fmt.Printf("==  Native  == : %s\n", native)
//	//
//	//Convey("The unmarshalled data should match the raw data JSON data", t, func() {
//	//	So(string(processed), ShouldEqual, string(native))
//	//})
//}

func TestAddHeaders(t *testing.T) {
	c := Client{}

	tests := []struct {
		method        string
		expectAuth    bool
		expectAccept  bool
		expectContent bool
	}{
		{
			method:       http.MethodGet,
			expectAuth:   true,
			expectAccept: true,
		},
		{
			method:        http.MethodPost,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:        http.MethodPut,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:        http.MethodPatch,
			expectAuth:    true,
			expectAccept:  true,
			expectContent: true,
		},
		{
			method:     http.MethodDelete,
			expectAuth: true,
		},
	}

	Convey("Given a http method", t, func() {
		for _, test := range tests {
			Convey(fmt.Sprintf("When the addHeaders method is called with the %s method in the request", test.method), func() {
				req, _ := http.NewRequest(test.method, "http://no-where", nil)
				c.addHeaders(req)

				Convey("Then the expected headers are added to the request", func() {
					So(headerExists(req, "accept"), ShouldEqual, test.expectAccept)
					So(headerExists(req, "content-type"), ShouldEqual, test.expectContent)
					So(headerExists(req, "Authorization"), ShouldEqual, test.expectAuth)
				})
			})
		}
	})
}
