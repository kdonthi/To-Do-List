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
	Items []ItemAndID `json:"items"`
}

type ItemAndID struct {
	Item string `json:"item"`
	ID   int    `json:"ID"`
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
			itemList.CreateItem(reqBody.Item)
			writer.Write(responseBody{item: b})
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

		item, err := itemList.ReadItem(index)
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			b, err := json.Marshal(item)
			if err != nil {
				writer.Write([]byte(err.Error()))
			}

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

		err = itemList.UpdateItem(index, b)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		writer.Write(b)
	})
}

func DeleteItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		deletedItem, err := itemList.DeleteItem(index)
		if err != nil {
			writer.Write([]byte(err.Error()))
		}

		writer.Write([]byte(deletedItem))
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
