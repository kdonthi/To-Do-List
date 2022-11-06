package backend

import (
	"TodoApplication/utils"
	"github.com/julienschmidt/httprouter"
)

func SetHandlers(router *httprouter.Router, itemList *utils.ItemList) {
	router.GET("/", PrintItems(itemList))
	router.POST("/create", CreateItem(itemList))
	router.GET("/read/:id", ReadItem(itemList))
	router.GET("/read", ReadAll(itemList))
	router.PUT("/update/:id", UpdateItem(itemList))
	router.DELETE("/delete/:id", DeleteItem(itemList))
	router.DELETE("/delete", DeleteAll(itemList))
}
