package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestListItems(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		expectedResponse string
	}{
		{
			name:             "no items",
			itemList:         NewItemList(),
			expectedResponse: "",
		},
		{
			name:             "one item",
			itemList:         &ItemList{items: []string{"abc"}},
			expectedResponse: "1. abc\n",
		},
		{
			name:             "multiple items",
			itemList:         &ItemList{items: []string{"abc", "{hello:world}", "123"}},
			expectedResponse: "1. abc\n2. {hello:world}\n3. 123\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			testCase.itemList.ListItems(recorder)

			b, err := io.ReadAll(recorder.Body)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedResponse, string(b))
		})
	}
}

func TestCreateItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		itemToAdd        string
		expectedItemList *ItemList
	}{
		{
			name:             "no existing items",
			itemList:         NewItemList(),
			itemToAdd:        "hello",
			expectedItemList: &ItemList{items: []string{"hello"}},
		},
		{
			name:             "existing item",
			itemList:         &ItemList{items: []string{"hello"}},
			itemToAdd:        "world",
			expectedItemList: &ItemList{items: []string{"hello", "world"}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.itemList.CreateItem(testCase.itemToAdd)
			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
		})
	}
}

func TestReadItem(t *testing.T) {
	testTable := []struct {
		name          string
		itemList      *ItemList
		index         int
		expectedItem  string
		expectedError error
	}{
		{
			name:          "index is 0",
			itemList:      &ItemList{items: []string{"abc", "bcd"}},
			index:         0,
			expectedItem:  "",
			expectedError: fmt.Errorf("item number less than 1"),
		},
		{
			name:          "index is more than length",
			itemList:      &ItemList{items: []string{"abc", "bcd"}},
			index:         3,
			expectedItem:  "",
			expectedError: fmt.Errorf("item number %v more than number of items %v", 3, 2),
		},
		{
			name:          "index is 1",
			itemList:      &ItemList{items: []string{"abc", "bcd"}},
			index:         1,
			expectedItem:  "abc",
			expectedError: nil,
		},
		{
			name:          "index is length",
			itemList:      &ItemList{items: []string{"abc", "bcd"}},
			index:         2,
			expectedItem:  "bcd",
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			item, err := testCase.itemList.ReadItem(testCase.index)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedItem, item)
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
		expectedError    error
	}{
		{
			name:             "index is 0",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            0,
			update:           "",
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedError:    fmt.Errorf("item number less than 1"),
		},
		{
			name:             "index is more than length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            3,
			update:           "",
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedError:    fmt.Errorf("item number %v more than number of items %v", 3, 2),
		},
		{
			name:             "index is 1",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            1,
			update:           "123",
			expectedItemList: &ItemList{items: []string{"123", "bcd"}},
			expectedError:    nil,
		},
		{
			name:             "index is length",
			itemList:         &ItemList{items: []string{"123", "bcd"}},
			index:            2,
			update:           "456",
			expectedItemList: &ItemList{items: []string{"123", "456"}},
			expectedError:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.itemList.UpdateItem(testCase.index, testCase.update)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
		})
	}
}

func TestDeleteItem(t *testing.T) {
	testTable := []struct {
		name             string
		itemList         *ItemList
		index            int
		expectedItemList *ItemList
		expectedError    error
	}{
		{
			name:             "index is 0",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            0,
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedError:    fmt.Errorf("item number less than 1"),
		},
		{
			name:             "index is more than length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            3,
			expectedItemList: &ItemList{items: []string{"abc", "bcd"}},
			expectedError:    fmt.Errorf("item number %v more than number of items %v", 3, 2),
		},
		{
			name:             "index is 1",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            1,
			expectedItemList: &ItemList{items: []string{"bcd"}},
			expectedError:    nil,
		},
		{
			name:             "index is length",
			itemList:         &ItemList{items: []string{"abc", "bcd"}},
			index:            2,
			expectedItemList: &ItemList{items: []string{"abc"}},
			expectedError:    nil,
		},
		{
			name:             "index is in middle",
			itemList:         &ItemList{items: []string{"abc", "bcd", "cdf", "123"}},
			index:            3,
			expectedItemList: &ItemList{items: []string{"abc", "bcd", "123"}},
			expectedError:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.itemList.DeleteItem(testCase.index)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedItemList, testCase.itemList)
		})
	}
}
