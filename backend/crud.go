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

type RequestBody struct {
	Item string `json:"item"`
}

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

func Count(itemList *utils.ItemList) httprouter.Handle {
	return httprouter.Handle(func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		count := len(itemList.ReadAll())
		writer.Write([]byte(fmt.Sprintf("%v", count)))
	})
}

func getID(ps httprouter.Params) (int, error) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		return 0, fmt.Errorf("error converting id to number: %v", err)
	}

	return id, nil
}

func parseRequestBody(request *http.Request) (*RequestBody, error) {
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var r RequestBody
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	if r.Item == "" {
		return nil, fmt.Errorf("\"item\" field in body was not populated")
	}

	return &r, nil
}

func writeError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte(err.Error()))
}
