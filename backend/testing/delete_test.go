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

func TestDeleteItem_InvalidID(t *testing.T) {
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
			_, code, err := deleteItem(t, router, testCase.id)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, code)
		})
	}
}

func TestDeleteAll(t *testing.T) {
	router, itemList := setup()

	itemList.CreateItem("abc")
	itemList.CreateItem("def")

	deleteItems, code := deleteAll(t, router)
	require.Equal(t, []utils.ItemAndID{{ID: 1, Item: "abc"}, {ID: 2, Item: "def"}}, deleteItems)
	require.Equal(t, 200, code)

	assert.Equal(t, len(itemList.ReadAll()), 0)
}

func deleteItem(t *testing.T, router *httprouter.Router, id int) (utils.ItemAndID, int, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%v", id), nil)

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

func deleteAll(t *testing.T, router *httprouter.Router) ([]utils.ItemAndID, int) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/delete", nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp []utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp, w.Code
}
