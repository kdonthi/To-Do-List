package rpc

import "TodoApplication/utils"

type RequestBody struct {
	Item string `json:"item"`
}

type SingleResponseBody struct {
	Item utils.ItemAndID
}

type ListResponseBody struct {
	Items []utils.ItemAndID `json:"items"`
}
