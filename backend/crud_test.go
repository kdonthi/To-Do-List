package backend

import (
	"TodoApplication/utils"
	"bytes"
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

	b, err := listItems(router)
	require.Equal(t, "", string(b))

	b, err = createItem(t, router, "abc")
	require.Nil(t, err)
	require.Equal(t, "abc", string(b))

	b, err = readItem(router, 1)
	require.Nil(t, err)
	require.Equal(t, "abc", string(b))

	b, err = updateItem(router, 1, "123")
	require.Nil(t, err)
	require.Equal(t, "123", string(b))

	b, err = listItems(router)
	require.Equal(t, "1. 123\n", string(b))

	b, err = deleteItem(router, 1)
	require.Nil(t, err)
	require.Equal(t, "123", string(b))
}

//func TestValidIndex(t *testing.T) {
//	itemList := utils.NewItemList()
//	handle := CreateItem(itemList)
//}

func listItems(router *httprouter.Router) ([]byte, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	return io.ReadAll(w.Body)
}

func createItem(t *testing.T, router *httprouter.Router, item string) ([]byte, error) {
	w := httptest.NewRecorder()
	reqBody := bytes.NewBuffer([]byte(item))
	req, _ := http.NewRequest(http.MethodPost, "/create", reqBody)

	router.ServeHTTP(w, req)

	return io.ReadAll(w.Body)
}

func readItem(router *httprouter.Router, id int) ([]byte, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/read/%v", id), nil)

	router.ServeHTTP(w, req)

	return io.ReadAll(w.Body)
}

func updateItem(router *httprouter.Router, id int, newItem string) ([]byte, error) {
	w := httptest.NewRecorder()
	reqBody := bytes.NewBuffer([]byte(newItem))
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/update/%v", id), reqBody)

	router.ServeHTTP(w, req)

	return io.ReadAll(w.Body)
}

func deleteItem(router *httprouter.Router, id int) ([]byte, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%v", id), nil)

	router.ServeHTTP(w, req)

	return io.ReadAll(w.Body)
}
