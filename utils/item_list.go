package utils

import (
	"fmt"
)

type ItemList struct {
	items []string
}

type ItemAndID struct {
	Item string `json:"item"`
	ID   int    `json:"ID"`
}

type RequestBody struct {
	Item string `json:"item"`
}

type SingleResponseBody struct {
	Item ItemAndID `json:"item"`
}

type ListResponseBody struct {
	Items []ItemAndID `json:"items"`
}

func NewItemList() *ItemList {
	return &ItemList{
		items: []string{},
	}
}

func (il *ItemList) ListItems() []ItemAndID {
	var items []ItemAndID
	for i, item := range il.items {
		items = append(items, ItemAndID{
			Item: item,
			ID:   i + 1,
		})
	}

	return items
}

func (il *ItemList) CreateItem(item string) ItemAndID {
	il.items = append(il.items, item)

	return ItemAndID{
		Item: item,
		ID:   len(il.items),
	}
}

func (il *ItemList) ReadItem(index int) (ItemAndID, error) {
	err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	adjustedIndex := index - 1
	return ItemAndID{
		Item: il.items[adjustedIndex],
		ID:   index,
	}, nil
}

func (il *ItemList) DeleteItem(index int) (ItemAndID, error) {
	err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	adjustedIndex := index - 1
	itemToDelete := il.items[adjustedIndex]
	il.items = append(il.items[:adjustedIndex], il.items[adjustedIndex+1:]...)

	return ItemAndID{
		Item: itemToDelete,
		ID:   index,
	}, nil
}

func (il *ItemList) UpdateItem(index int, newItem string) (ItemAndID, error) {
	err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	adjustedIndex := index - 1
	il.items[adjustedIndex] = newItem

	return ItemAndID{
		Item: newItem,
		ID:   index,
	}, nil
}

func (il *ItemList) validateIndex(index int) error {
	if index < 1 {
		return fmt.Errorf("item number less than 1")
	} else if index > len(il.items) {
		return fmt.Errorf("item number (%v) is more than number of items (%v)", index, len(il.items))
	}
	return nil
}
