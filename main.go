package main

import (
	"TodoApplication/backend"
	"TodoApplication/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	itemList := utils.NewItemList()

	backend.SetHandlers(router, itemList)

	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatal(err)
	}
}
