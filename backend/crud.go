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

func PrintItems(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		items := itemList.ListItems()

		result := "TO-DO LIST\n" +
			"----------\n"

		for _, item := range items {
			result += fmt.Sprintf("%v. %v\n", item.ID, item.Item)
		}
		if len(items) == 0 {
			result += "Looking kind of empty...\n"
		}

		writer.Write([]byte(result))
	})
}

func ListItems(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		items := itemList.ListItems()

		b, err := json.Marshal(items)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		writer.Write(b)
	})

}

func CreateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		reqBody, err := getRequestBody(request)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		} else {
			item := itemList.CreateItem(reqBody.Item)

			b, err := json.Marshal(item)
			if err != nil {
				writer.Write([]byte(err.Error()))
				return
			}

			writer.Write(b)
		}
	})
}

func ReadItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		id, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		item, err := itemList.ReadItem(id)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		} else {
			b, err := json.Marshal(item)
			if err != nil {
				writer.Write([]byte(err.Error()))
				return
			}

			writer.Write(b)
		}
	})
}

func UpdateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		id, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		reqBody, err := getRequestBody(request)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		item, err := itemList.UpdateItem(id, reqBody.Item)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		b, err := json.Marshal(item)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		writer.Write(b)
	})
}

func DeleteItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		item, err := itemList.DeleteItem(index)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		b, err := json.Marshal(item)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

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

func getRequestBody(request *http.Request) (*utils.RequestBody, error) {
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var r utils.RequestBody
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
