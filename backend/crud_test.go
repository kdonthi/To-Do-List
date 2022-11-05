package backend

import (
	"TodoApplication/utils"
	"bytes"
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

func setup() *httprouter.Router {
	router := httprouter.New()
	itemList := utils.NewItemList()
	SetHandlers(router, itemList)
	return router
}

func TestCrud(t *testing.T) {
	router := setup()

	itemsStr := printItems(t, router)
	require.Equal(t, "TO-DO LIST\n----------\nLooking kind of empty...\n", itemsStr)

	items := listItems(t, router)
	require.Equal(t, []utils.ItemAndID{}, items)

	item := createItem(t, router, "abc")
	require.Equal(t, utils.ItemAndID{
		Item: "abc",
		ID:   1,
	}, item)

	item, err := readItem(t, router, 1)
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{
		Item: "abc",
		ID:   1,
	}, item)

	item, err = updateItem(t, router, 1, "123")
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{
		Item: "123",
		ID:   1,
	}, item)

	itemsStr = printItems(t, router)
	require.Equal(t, "TO-DO LIST\n----------\n1. 123\n", itemsStr)

	items = listItems(t, router)
	require.Equal(t, []utils.ItemAndID{
		{
			Item: "123",
			ID:   1,
		},
	}, items)

	item, err = deleteItem(t, router, 1)
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{
		Item: "123",
		ID:   1,
	}, item)
}

func TestReadItem_InvalidID(t *testing.T) {
	testTable := []struct {
		name          string
		id            int
		expectedError error
	}{
		{
			name:          "valid id",
			id:            2,
			expectedError: nil,
		},
		{
			name:          "invalid id below 1",
			id:            0,
			expectedError: fmt.Errorf("id is less than 1"),
		},
		{
			name:          "invalid id above size",
			id:            3,
			expectedError: fmt.Errorf("id (3) is more than the number of items (2)"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router := setup()

			createItem(t, router, "hello")
			createItem(t, router, "world")

			_, err := readItem(t, router, testCase.id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestUpdateItem_InvalidID(t *testing.T) {
	testTable := []struct {
		name          string
		id            int
		expectedError error
	}{
		{
			name:          "valid id",
			id:            2,
			expectedError: nil,
		},
		{
			name:          "invalid id below 1",
			id:            0,
			expectedError: fmt.Errorf("id is less than 1"),
		},
		{
			name:          "invalid id above size",
			id:            3,
			expectedError: fmt.Errorf("id (3) is more than the number of items (2)"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router := setup()

			createItem(t, router, "hello")
			createItem(t, router, "world")

			_, err := updateItem(t, router, testCase.id, "abc")
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestDeleteItem_InvalidID(t *testing.T) {
	testTable := []struct {
		name          string
		id            int
		expectedError error
	}{
		{
			name:          "valid id",
			id:            2,
			expectedError: nil,
		},
		{
			name:          "invalid id below 1",
			id:            0,
			expectedError: fmt.Errorf("id is less than 1"),
		},
		{
			name:          "invalid id above size",
			id:            3,
			expectedError: fmt.Errorf("id (3) is more than the number of items (2)"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			router := setup()

			createItem(t, router, "hello")
			createItem(t, router, "world")

			_, err := deleteItem(t, router, testCase.id)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func printItems(t *testing.T, router *httprouter.Router) string {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	return string(b)
}

func listItems(t *testing.T, router *httprouter.Router) []utils.ItemAndID {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/read-all", nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp []utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp
}

func createItem(t *testing.T, router *httprouter.Router, item string) utils.ItemAndID {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/create", reqBody(item))

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp
}

func readItem(t *testing.T, router *httprouter.Router, id int) (utils.ItemAndID, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/read/%v", id), nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return utils.ItemAndID{}, fmt.Errorf(string(b))
	}

	return resp, nil
}

func updateItem(t *testing.T, router *httprouter.Router, id int, newItem string) (utils.ItemAndID, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/update/%v", id), reqBody(newItem))

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return utils.ItemAndID{}, fmt.Errorf(string(b))
	}

	return resp, nil
}

func deleteItem(t *testing.T, router *httprouter.Router, id int) (utils.ItemAndID, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%v", id), nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return utils.ItemAndID{}, fmt.Errorf(string(b))
	}

	return resp, nil
}

func reqBody(item string) *bytes.Buffer {
	b, _ := json.Marshal(utils.RequestBody{Item: item})
	return bytes.NewBuffer(b)
}
