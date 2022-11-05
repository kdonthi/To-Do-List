package backend

import (
	"TodoApplication/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

type requestBody struct {
	Item string `json:"item"`
}

type responseBody struct {
	Items []utils.ItemAndID `json:"items"`
}

func ListItems(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		itemList.ListItems(writer)
	})
}

func CreateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		reqBody, err := getRequestBody(request)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			itemAndID := itemList.CreateItem(reqBody.Item)
			b, _ := json.Marshal(itemAndID)
			writer.Write(b)
		}
	})
}

func getRequestBody(request *http.Request) (*requestBody, error) {
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var r requestBody
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func ReadItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		itemAndID, err := itemList.ReadItem(index)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			b, _ := json.Marshal(itemAndID)
			writer.Write(b)
		}
	})
}

func UpdateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		b, err := io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		itemAndID, err := itemList.UpdateItem(index, string(b))
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		b, _ = json.Marshal(itemAndID)
		writer.Write(b)
	})
}

func DeleteItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		itemAndID, err := itemList.DeleteItem(index)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		b, _ := json.Marshal(itemAndID)
		writer.Write(b)
	})
}

func getID(ps httprouter.Params) (int, error) {
	idVar := ps.ByName("id")
	if idVar == "" {
		return 0, fmt.Errorf("id url parameter not set (usage: /endpoint/{id})")
	}

	index, err := strconv.Atoi(idVar)
	if err != nil {
		return 0, fmt.Errorf("error converting id to number: %v", err)
	}

	return index, nil
}
