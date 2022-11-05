package backend

import (
	"TodoApplication/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrud(t *testing.T) {
	router := httprouter.New()
	itemList := utils.NewItemList()

	SetHandlers(router, itemList)

	items := listItems(t, router)
	require.Equal(t, utils.ListResponseBody{}, items)

	item := createItem(t, router, "abc")
	require.Equal(t, utils.ItemAndID{
		Item: "abc",
		ID:   1,
	}, item)

	item = readItem(t, router, 1)
	require.Equal(t, utils.ItemAndID{
		Item: "abc",
		ID:   1,
	}, item)

	item = updateItem(t, router, 1, "123")
	require.Equal(t, utils.ItemAndID{
		Item: "123",
		ID:   1,
	}, item)

	items = listItems(t, router)
	require.Equal(t, utils.ListResponseBody{Items: []utils.ItemAndID{
		{
			Item: "123",
			ID:   1,
		},
	}}, items)

	item = deleteItem(t, router, 1)
	require.Equal(t, utils.ItemAndID{
		Item: "123",
		ID:   1,
	}, item)
}

func listItems(t *testing.T, router *httprouter.Router) utils.ListResponseBody {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ListResponseBody
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

func readItem(t *testing.T, router *httprouter.Router, id int) utils.ItemAndID {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/read/%v", id), nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp
}

func updateItem(t *testing.T, router *httprouter.Router, id int, newItem string) utils.ItemAndID {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/update/%v", id), reqBody(newItem))

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp
}

func deleteItem(t *testing.T, router *httprouter.Router, id int) utils.ItemAndID {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%v", id), nil)

	router.ServeHTTP(w, req)

	b, err := io.ReadAll(w.Body)
	require.Nil(t, err)

	var resp utils.ItemAndID
	require.Nil(t, json.Unmarshal(b, &resp))

	return resp
}

func reqBody(item string) *bytes.Buffer {
	b, _ := json.Marshal(RequestBody{Item: item})
	return bytes.NewBuffer(b)
}
