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

func TestUpdateItem_IDValidity(t *testing.T) {
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
			_, code, err := updateItemValidBody(t, router, testCase.id, "abc")

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, code)
		})
	}
}

func TestUpdateItem_InvalidBody(t *testing.T) {
	router, itemList := setup()

	itemList.CreateItem("hello")
	_, code, err := updateItemInvalidBody(t, router, 1)

	require.NotNil(t, err)
	assert.Equal(t, "item field in body was not populated", err.Error())
	assert.Equal(t, 400, code)
}

func updateItemValidBody(t *testing.T, router *httprouter.Router, id int, newItem string) (utils.ItemAndID, int, error) {
	return updateItem(t, router, id, createValidRequestBody(newItem))
}

func updateItemInvalidBody(t *testing.T, router *httprouter.Router, id int) (utils.ItemAndID, int, error) {
	return updateItem(t, router, id, createInvalidRequestBody())
}

func updateItem(t *testing.T, router *httprouter.Router, id int, body io.Reader) (utils.ItemAndID, int, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/update/%v", id), body)

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
