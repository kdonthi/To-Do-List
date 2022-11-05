package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListItems(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		expectedResponse []ItemAndID
	}{
		{
			name:             "no items",
			itemList:         NewItemList(),
			expectedResponse: nil,
		},
		{
			name:     "one item",
			itemList: &ItemList{items: []string{"abc"}},
			expectedResponse: []ItemAndID{
				{
					Item: "abc",
					ID:   1,
				},
			},
		},
		{
			name:     "multiple items",
			itemList: &ItemList{items: []string{"abc", "{hello:world}", "123"}},
			expectedResponse: []ItemAndID{
				{
					Item: "abc",
					ID:   1,
				},
				{
					Item: "{hello:world}",
					ID:   2,
				},
				{
					Item: "123",
					ID:   3,
				},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			items := testCase.itemList.ListItems()

			assert.Equal(t, testCase.expectedResponse, items)
		})
	}
}

func TestCreateItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		itemToAdd        string
		expectedItemList *ItemList
		expectedResponse ItemAndID
		expectedError    error
	}{
		{
			name:             "no existing items",
			itemList:         NewItemList(),
			itemToAdd:        "hello",
			expectedItemList: &ItemList{items: []string{"hello"}},
			expectedResponse: ItemAndID{
				Item: "hello",
				ID:   1,
			},
		},
		{
			name:             "existing item",
			itemList:         &ItemList{items: []string{"hello"}},
			itemToAdd:        "world",
			expectedItemList: &ItemList{items: []string{"hello", "world"}},
			expectedResponse: ItemAndID{
				Item: "world",
				ID:   2,
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			item := testCase.itemList.CreateItem(testCase.itemToAdd)

			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
			assert.Equal(t, testCase.expectedResponse, item)
		})
	}
}

func TestReadItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		index            int
		expectedResponse ItemAndID
		expectedError    error
	}{
		{
			name:             "id is 0",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            0,
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number less than 1"),
		},
		{
			name:             "id is more than length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            3,
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number (%v) is more than number of items (%v)", 3, 2),
		},
		{
			name:     "id is 1",
			itemList: &ItemList{items: []string{"abc", "bcd"}},
			index:    1,
			expectedResponse: ItemAndID{
				Item: "abc",
				ID:   1,
			},
			expectedError: nil,
		},
		{
			name:     "id is length",
			itemList: &ItemList{items: []string{"abc", "bcd"}},
			index:    2,
			expectedResponse: ItemAndID{
				Item: "bcd",
				ID:   2,
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			item, err := testCase.itemList.ReadItem(testCase.index)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResponse, item)
		})
	}
}

func TestUpdateItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		index            int
		update           string
		expectedItemList *ItemList
		expectedResponse ItemAndID
		expectedError    error
	}{
		{
			name:             "id is 0",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            0,
			update:           "",
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number less than 1"),
		},
		{
			name:             "id is more than length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            3,
			update:           "",
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number (%v) is more than number of items (%v)", 3, 2),
		},
		{
			name:             "id is 1",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            1,
			update:           "123",
			expectedItemList: &ItemList{items: []string{"123", "bcd"}},
			expectedResponse: ItemAndID{
				Item: "123",
				ID:   1,
			},
			expectedError: nil,
		},
		{
			name:             "id is length",
			itemList:         &ItemList{items: []string{"123", "bcd"}},
			index:            2,
			update:           "456",
			expectedItemList: &ItemList{items: []string{"123", "456"}},
			expectedResponse: ItemAndID{
				Item: "456",
				ID:   2,
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			item, err := testCase.itemList.UpdateItem(testCase.index, testCase.update)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
			assert.Equal(t, testCase.expectedResponse, item)
		})
	}
}

func TestDeleteItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		id               int
		expectedItemList *ItemList
		expectedResponse ItemAndID
		expectedError    error
	}{
		{
			name:             "id is 0",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			id:               0,
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number less than 1"),
		},
		{
			name:             "id is more than length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			id:               3,
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedResponse: ItemAndID{},
			expectedError:    fmt.Errorf("item number (%v) is more than number of items (%v)", 3, 2),
		},
		{
			name:             "id is 1",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			id:               1,
			expectedItemList: &ItemList{items: []string{"bcd"}},
			expectedResponse: ItemAndID{
				Item: "abc",
				ID:   1,
			},
			expectedError: nil,
		},
		{
			name:             "id is length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			id:               2,
			expectedItemList: &ItemList{items: []string{"abc"}},
			expectedResponse: ItemAndID{
				Item: "bcd",
				ID:   2,
			},
			expectedError: nil,
		},
		{
			name:             "id is in middle",
			itemList:         &ItemList{items: []string{"abc", "bcd", "cdf", "123"}},
			id:               3,
			expectedItemList: &ItemList{items: []string{"abc", "bcd", "123"}},
			expectedResponse: ItemAndID{
				Item: "cdf",
				ID:   3,
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			item, err := testCase.itemList.DeleteItem(testCase.id)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
			assert.Equal(t, testCase.expectedResponse, item)
		})
	}
}

func TestDeleteAll(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		expectedItemList *ItemList
		expectedResponse []ItemAndID
	}{
		{
			name:             "0 items",
			itemList:         &ItemList{items: []string{}},
			expectedItemList: &ItemList{items: []string{}},
			expectedResponse: nil,
		},
		{
			name:             "2 items",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			expectedItemList: &ItemList{items: []string{}},
			expectedResponse: []ItemAndID{
				{
					ID:   1,
					Item: "abc",
				},
				{
					ID:   2,
					Item: "bcd",
				},
			},
		},
		{
			name:             "4 items",
			itemList:         &ItemList{items: []string{"abc", "bcd", "cdf", "123"}},
			expectedItemList: &ItemList{items: []string{}},
			expectedResponse: []ItemAndID{
				{
					ID:   1,
					Item: "abc",
				},
				{
					ID:   2,
					Item: "bcd",
				},
				{
					ID:   3,
					Item: "cdf",
				},
				{
					ID:   4,
					Item: "123",
				},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			items := testCase.itemList.DeleteAll()

			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
			assert.Equal(t, testCase.expectedResponse, items)
		})
	}
}
