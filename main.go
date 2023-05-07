package main

import (
	"fmt"
	"hotelmenu/order"
)

func main() {

	fileName := "log.txt"

	sortedMenuIDs, orderedItemCounts, err := order.FetchTopThreeOrderedItems(fileName)
	if err != nil {
		panic(err)
	}

	// print the top 3 menu items consumed
	fmt.Println("Top 3 menu items consumed:")
	for i, menuID := range sortedMenuIDs {
		if i >= 3 {
			break
		}
		fmt.Printf("%d. %d (%d)\n", i+1, menuID, orderedItemCounts[menuID])
	}
}
