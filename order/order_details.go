package order

import (
	"bufio"
	"fmt"
	"hotelmenu/errors"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// declared error constants
const (
	InvalidFileErr errors.Code = iota + 700
	FileNotFoundErr
	IncorrectInputErr
)

type order struct {
	eaterID    int
	foodMenuID int
}

func FetchTopThreeOrderedItems(fileName string) (sortedMenuIDs []int, orderedItemCounts map[int]int, err error) {
	// Extract the file extension using the filepath package
	ext := filepath.Ext(fileName)

	// Check if the extension is valid
	if strings.ToLower(ext) != ".txt" {
		err = errors.NewUnknownError(InvalidFileErr, fmt.Errorf("Invalid file type1."))
		return
	}

	orderDetails, err := fetchOrderDetails(fileName)
	if err != nil {
		return
	}

	orderedItemCounts, err = fetchCountOfOrderedItemByEater(orderDetails)
	if err != nil {
		return
	}

	sortedMenuIDs = sortItemsByCountDesc(orderedItemCounts)

	return
}

func fetchOrderDetails(filename string) (orders []order, err error) {
	// Open the log file
	file, err := os.Open(filename)
	if err != nil {
		err = errors.NewNotFoundError(FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		itemDetails := scanner.Text()

		// Split the line into the eater_id and foodmenu_id
		parts := strings.Split(itemDetails, ",")
		if len(parts) != 2 {
			err = errors.NewIncorrectInputError(IncorrectInputErr, fmt.Errorf("Invalid order log entry"))
			return
		}

		eaterID, err := strconv.Atoi(strings.Trim(parts[0], " "))
		if err != nil {
			err = errors.NewIncorrectInputError(IncorrectInputErr, fmt.Errorf("Invalid eater_id: %s", parts[0]))
			return nil, err
		}

		foodMenuID, err := strconv.Atoi(strings.Trim(parts[1], " "))
		if err != nil {
			err = errors.NewIncorrectInputError(IncorrectInputErr, fmt.Errorf("Invalid foodmenu_id: %s", parts[1]))
			return nil, err
		}

		orders = append(orders, order{eaterID, foodMenuID})
	}
	return
}

func fetchCountOfOrderedItemByEater(orderDetails []order) (map[int]int, error) {
	var orderedItemCount = make(map[int]int)
	var err error

	// create a map to keep track of the eater_id and the foodmenu_id consumed
	consumed := make(map[int]int)
	for _, orderedItem := range orderDetails {

		// check if this eater_id has already consumed this foodmenu_id
		if prevFoodmenuID, ok := consumed[orderedItem.eaterID]; ok {
			if prevFoodmenuID == orderedItem.foodMenuID {
				err = errors.NewIncorrectInputError(IncorrectInputErr, fmt.Errorf("Duplicate entry found for eater_id=%d and foodmenu_id=%d\n", orderedItem.eaterID, orderedItem.foodMenuID))
				return nil, err
			}
		}

		// Increment the count for this food item
		orderedItemCount[orderedItem.foodMenuID]++
		consumed[orderedItem.eaterID] = orderedItem.foodMenuID
	}
	return orderedItemCount, err
}

func sortItemsByCountDesc(orderedItemCount map[int]int) []int {

	// sort the foodmenu_ids by their consumption count
	sortedMenuIDs := make([]int, 0, len(orderedItemCount))
	for menuID := range orderedItemCount {
		sortedMenuIDs = append(sortedMenuIDs, menuID)
	}
	sort.Slice(sortedMenuIDs, func(i, j int) bool {
		return orderedItemCount[sortedMenuIDs[i]] > orderedItemCount[sortedMenuIDs[j]]
	})
	return sortedMenuIDs
}
