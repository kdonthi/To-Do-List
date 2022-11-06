package testing

import (
	"TodoApplication/backend"
	"TodoApplication/utils"
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func setup() (*httprouter.Router, *utils.ItemList) {
	router := httprouter.New()
	itemList := utils.NewItemList()
	backend.SetHandlers(router, itemList)
	return router, itemList
}

func createValidRequestBody(item string) *bytes.Buffer {
	b, _ := json.Marshal(backend.RequestBody{Item: item})
	return bytes.NewBuffer(b)
}

type invalidRequestBody struct {
	Noop string `json:"noop"`
}

func createInvalidRequestBody() *bytes.Buffer {
	b, _ := json.Marshal(invalidRequestBody{Noop: "123"})
	return bytes.NewBuffer(b)
}
