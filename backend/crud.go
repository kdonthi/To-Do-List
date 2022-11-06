package backend

import (
	"TodoApplication/rpc"
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
		items := itemList.ReadAll()

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

func CreateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		reqBody, err := parseRequestBody(request)
		if err != nil {
			writeError(writer, err)
			return
		} else {
			item := itemList.CreateItem(reqBody.Item)

			b, err := json.Marshal(item)
			if err != nil {
				writeError(writer, err)
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
			writeError(writer, err)
			return
		}

		item, err := itemList.ReadItem(id)
		if err != nil {
			writeError(writer, err)
			return
		} else {
			b, err := json.Marshal(item)
			if err != nil {
				writeError(writer, err)
				return
			}

			writer.Write(b)
		}
	})
}

func ReadAll(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		items := itemList.ReadAll()

		b, err := json.Marshal(items)
		if err != nil {
			writeError(writer, err)
			return
		}

		writer.Write(b)
	})
}

func UpdateItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		id, err := getID(ps)
		if err != nil {
			writeError(writer, err)
			return
		}

		reqBody, err := parseRequestBody(request)
		if err != nil {
			writeError(writer, err)
			return
		}

		if reqBody.Item == "" {
			writeError(writer, err)
			return
		}

		item, err := itemList.UpdateItem(id, reqBody.Item)
		if err != nil {
			writeError(writer, err)
			return
		}

		b, err := json.Marshal(item)
		if err != nil {
			writeError(writer, err)
			return
		}

		writer.Write(b)
	})
}

func DeleteItem(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		index, err := getID(ps)
		if err != nil {
			writeError(writer, err)
			return
		}

		item, err := itemList.DeleteItem(index)
		if err != nil {
			writeError(writer, err)
			return
		}

		b, err := json.Marshal(item)
		if err != nil {
			writeError(writer, err)
			return
		}

		writer.Write(b)
	})
}

func DeleteAll(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		items := itemList.DeleteAll()

		b, err := json.Marshal(items)
		if err != nil {
			writeError(writer, err)
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

func parseRequestBody(request *http.Request) (*rpc.RequestBody, error) {
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var r rpc.RequestBody
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if r.Item == "" {
		return nil, fmt.Errorf("item field in body was not populated")
	}

	return &r, nil
}

func writeError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte(err.Error()))
}
