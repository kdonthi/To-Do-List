package testing

import (
	"TodoApplication/backend"
	"TodoApplication/utils"
	"github.com/julienschmidt/httprouter"
)

func setup() (*httprouter.Router, *utils.ItemList) {
	router := httprouter.New()
	itemList := utils.NewItemList()
	backend.SetHandlers(router, itemList)
	return router, itemList
}
