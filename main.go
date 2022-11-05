package main

import (
	"TodoApplication/backend"
	"TodoApplication/utils"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := httprouter.New()
	itemList := utils.NewItemList()
	port := 9000

	if len(os.Args) >= 2 {
		newPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("argument for port (%v) has to be an integer", os.Args[1])
		}

		port = newPort
	}

	backend.SetHandlers(router, itemList)

	logrus.Infof("server starting at port %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
