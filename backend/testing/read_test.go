package testing

import (
	"TodoApplication/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrintItems(t *testing.T) {
	testTable := []struct {
		name             string
		values           []string
		expectedResponse string
	}{
		{
			name:             "no items",
			values:           []string{},
			expectedResponse: "TO-DO LIST\n----------\nLooking kind of empty...\n",
		},
		{
			name:             "one item",
			values:           []string{"abc"},
			expectedResponse: "TO-DO LIST\n----------\n1. abc\n",
		},
		{
			name:             "multiple items",
			values:           []string{"abc", "def", "123"},
			expectedResponse: "TO-DO LIST\n----------\n1. abc\n2. def\n3. 123\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router, itemList := setup()

			for _, val := range testCase.values {
				itemList.CreateItem(val)
			}

			resp, code := printItems(t, router)
			assert.Equal(t, testCase.expectedResponse, resp)
			assert.Equal(t, 200, code)
		})
	}
}

func TestReadItem_IDValidity(t *testing.T) {
	testTable := []struct {
		name          string
		id            int
		expectedError error
		expectedCode  int
	}{
		{
			name:          "valid id",
			id:            2,
			expectedError: nil,
			expectedCode:  200,
		},
		{
			name:          "invalid id below 1",
			id:            0,
			expectedError: fmt.Errorf("id is less than 1"),
			expectedCode:  400,
		},
		{
			name:          "invalid id above size",
			id:            3,
			expectedError: fmt.Errorf("id (3) is more than the number of items (2)"),
			expectedCode:  400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router, itemList := setup()

			itemList.CreateItem("hello")
			itemList.CreateItem("world")
			_, code, err := readItem(t, router, testCase.id)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, code)
		})
	}
}

func TestReadAll(t *testing.T) {
	testTable := []struct {
		name             string
		values           []string
		expectedResponse []utils.ItemAndID
	}{
		{
			name:             "no items",
			values:           []string{},
			expectedResponse: []utils.ItemAndID{},
		},
		{
			name:             "one item",
			values:           []string{"abc"},
			expectedResponse: []utils.ItemAndID{{ID: 1, Item: "abc"}},
		},
		{
			name:             "multiple items",
			values:           []string{"abc", "def", "123"},
			expectedResponse: []utils.ItemAndID{{ID: 1, Item: "abc"}, {ID: 2, Item: "def"}, {ID: 3, Item: "123"}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router, itemList := setup()

			for _, val := range testCase.values {
				itemList.CreateItem(val)
			}

			resp, code := readItems(t, router)
			assert.Equal(t, testCase.expectedResponse, resp)
			assert.Equal(t, 200, code)
		})
	}
}

func TestCount(t *testing.T) {
	testTable := []struct {
		name          string
		values        []string
		expectedCount string
	}{
		{
			name:          "no items",
			values:        []string{},
			expectedCount: "0",
		},
		{
			name:          "one item",
			values:        []string{"abc"},
			expectedCount: "1",
		},
		{
			name:          "multiple items",
			values:        []string{"abc", "def", "123"},
			expectedCount: "3",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router, itemList := setup()

			for _, val := range testCase.values {
				itemList.CreateItem(val)
			}

			response, code := count(t, router)
			assert.Equal(t, testCase.expectedCount, response)
			assert.Equal(t, 200, code)
		})
	}
}

func printItems(t *testing.T, router *httprouter.Router) (string, int) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)
	code := w.Code

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	return string(b), code
}

func readItem(t *testing.T, router *httprouter.Router, id int) (utils.ItemAndID, int, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/read/%v", id), nil)

	router.ServeHTTP(w, req)
	code := w.Code

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return utils.ItemAndID{}, code, fmt.Errorf(string(b))
	}

	return resp, code, nil
}

func readItems(t *testing.T, router *httprouter.Router) ([]utils.ItemAndID, int) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/read", nil)

	router.ServeHTTP(w, req)
	code := w.Code

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp []utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp, code
}

func count(t *testing.T, router *httprouter.Router) (string, int) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/read", nil)

	router.ServeHTTP(w, req)
	code := w.Code

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	return string(b), code
}
