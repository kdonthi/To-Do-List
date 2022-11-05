package utils

import (
	"fmt"
)

type ItemList struct {
	items []string
}

type ItemAndID struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func NewItemList() *ItemList {
	return &ItemList{
		items: []string{},
	}
}

func (il *ItemList) ListItems() []ItemAndID {
	return itemsWithID(il.items)
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
	adjustedIndex, err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	itemToDelete := il.items[adjustedIndex]
	il.items = append(il.items[:adjustedIndex], il.items[adjustedIndex+1:]...)

	return ItemAndID{
		Item: itemToDelete,
		ID:   index,
	}, nil
}

func (il *ItemList) UpdateItem(index int, newItem string) (ItemAndID, error) {
	adjustedIndex, err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	il.items[adjustedIndex] = newItem

	return ItemAndID{
		Item: newItem,
		ID:   index,
	}, nil
}

func (il *ItemList) validateIndex(index int) error {
	if index < 1 {
		return fmt.Errorf("id is less than 1")
	} else if index > len(il.items) {
		return fmt.Errorf("id (%v) is more than the number of items (%v)", index, len(il.items))
	}
	return nil
}

func itemsWithID(listCpy []string) []ItemAndID {
	var items []ItemAndID
	for i, item := range listCpy {
		items = append(items, ItemAndID{
			ID:   i + 1,
			Item: item,
		})
	}

	return items
}
