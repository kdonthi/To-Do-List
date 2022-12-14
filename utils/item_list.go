package utils

import (
	"fmt"
	"sync"
)

type ItemList struct {
	items []string
	m     sync.RWMutex
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

func (il *ItemList) CreateItem(item string) ItemAndID {
	il.m.Lock()
	defer il.m.Unlock()

	il.items = append(il.items, item)

	return ItemAndID{
		Item: item,
		ID:   len(il.items),
	}
}

func (il *ItemList) ReadItem(index int) (ItemAndID, error) {
	adjustedIndex, err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	il.m.RLock()
	defer il.m.RUnlock()

	item := il.items[adjustedIndex]

	return ItemAndID{
		Item: item,
		ID:   index,
	}, nil
}

func (il *ItemList) ReadAll() []ItemAndID {
	il.m.Lock()
	defer il.m.Unlock()

	return itemsWithID(il.items)
}

func (il *ItemList) UpdateItem(index int, newItem string) (ItemAndID, error) {
	adjustedIndex, err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	il.m.Lock()
	defer il.m.Unlock()

	il.items[adjustedIndex] = newItem

	return ItemAndID{
		Item: newItem,
		ID:   index,
	}, nil
}

func (il *ItemList) DeleteItem(index int) (ItemAndID, error) {
	adjustedIndex, err := il.validateIndex(index)
	if err != nil {
		return ItemAndID{}, err
	}

	il.m.Lock()
	defer il.m.Unlock()

	itemToDelete := il.items[adjustedIndex]
	il.items = append(il.items[:adjustedIndex], il.items[adjustedIndex+1:]...)

	return ItemAndID{
		Item: itemToDelete,
		ID:   index,
	}, nil
}

func (il *ItemList) DeleteAll() []ItemAndID {
	listCpy := il.items

	il.m.Lock()
	defer il.m.Unlock()

	il.items = il.items[:0]
	return itemsWithID(listCpy)
}

func (il *ItemList) validateIndex(index int) (adjustedIndex int, err error) {
	if index < 1 {
		return 0, fmt.Errorf("id is less than 1")
	} else if index > len(il.items) {
		return 0, fmt.Errorf("id (%v) is more than the number of items (%v)", index, len(il.items))
	}
	return index - 1, nil
}

func itemsWithID(listCpy []string) []ItemAndID {
	items := []ItemAndID{}
	for i, item := range listCpy {
		items = append(items, ItemAndID{
			ID:   i + 1,
			Item: item,
		})
	}

	return items
}
