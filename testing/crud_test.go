package testing

import (
	"TodoApplication/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrud(t *testing.T) {
	router, _ := setup()

	itemsStr, _ := printItems(t, router)
	require.Equal(t, "TO-DO LIST\n----------\nLooking kind of empty...\n", itemsStr)

	items, _ := readItems(t, router)
	require.Equal(t, []utils.ItemAndID{}, items)

	item, _, _ := createItemValidBody(t, router, "abc")
	require.Equal(t, utils.ItemAndID{Item: "abc", ID: 1}, item)

	item, _, err := readItem(t, router, 1)
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{Item: "abc", ID: 1}, item)

	item, _, err = updateItemValidBody(t, router, 1, "123")
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{Item: "123", ID: 1}, item)

	itemsStr, _ = printItems(t, router)
	require.Equal(t, "TO-DO LIST\n----------\n1. 123\n", itemsStr)

	items, _ = readItems(t, router)
	require.Equal(t, []utils.ItemAndID{{Item: "123", ID: 1}}, items)

	item, _, err = deleteItem(t, router, 1)
	require.Nil(t, err)
	require.Equal(t, utils.ItemAndID{Item: "123", ID: 1}, item)

	items, _ = readItems(t, router)
	require.Equal(t, []utils.ItemAndID{}, items)
}
