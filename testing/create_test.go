package testing

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

func TestCreateItem_ValidBody(t *testing.T) {
	router, _ := setup()

	response, code, err := createItemValidBody(t, router, "abc")

	assert.Nil(t, err)
	assert.Equal(t, utils.ItemAndID{ID: 1, Item: "abc"}, response)
	assert.Equal(t, 200, code)
}

func TestCreateItem_InvalidBody(t *testing.T) {
	router, _ := setup()

	_, code, err := createItemInvalidBody(t, router)

	assert.Equal(t, "\"item\" field in body was not populated", err.Error())
	assert.Equal(t, 400, code)
}

func createItemValidBody(t *testing.T, router *httprouter.Router, item string) (utils.ItemAndID, int, error) {
	return createItem(t, router, createValidRequestBody(item))
}

func createItemInvalidBody(t *testing.T, router *httprouter.Router) (utils.ItemAndID, int, error) {
	return createItem(t, router, createInvalidRequestBody())
}

func createItem(t *testing.T, router *httprouter.Router, body *bytes.Buffer) (utils.ItemAndID, int, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/create", body)

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
