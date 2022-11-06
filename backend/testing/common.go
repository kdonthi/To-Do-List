package testing

import (
	"TodoApplication/backend"
	"TodoApplication/rpc"
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

func CreateValidRequestBody(item string) *bytes.Buffer {
	b, _ := json.Marshal(rpc.RequestBody{Item: item})
	return bytes.NewBuffer(b)
}

type InvalidRequestBody struct {
	Noop string `json:"noop"`
}

func CreateInvalidRequestBody() *bytes.Buffer {
	b, _ := json.Marshal(InvalidRequestBody{Noop: "123"})
	return bytes.NewBuffer(b)
}
