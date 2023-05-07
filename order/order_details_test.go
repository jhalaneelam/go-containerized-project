package order_test

import (
	"fmt"
	"hotelmenu/errors"
	"hotelmenu/order"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchTopThreeOrderedItems_FileError(t *testing.T) {

	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"invalidFile":  {fileName: "log", wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.InvalidFileErr},
		"fileNotFound": {fileName: "text.txt", wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.FileNotFoundErr},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			_, _, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
		})
	}
}

func TestFetchTopThreeOrderedItems_InvalidLogData(t *testing.T) {

	// Open the log file
	file, err := os.Open("log_test.txt")
	if err != nil {
		err = errors.NewNotFoundError(order.FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	os.WriteFile(file.Name(), []byte("1,1\n2 2\n"), 0664)
	defer file.Close()

	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"invalidOrderLogEntry": {fileName: file.Name(), wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.IncorrectInputErr},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			_, _, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
		})
	}
}

func TestFetchTopThreeOrderedItems_InvalidEaterId(t *testing.T) {

	// Open the log file
	file, err := os.Open("log_test.txt")
	if err != nil {
		err = errors.NewNotFoundError(order.FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	os.WriteFile(file.Name(), []byte("A,1\n2 2\n"), 0664)
	defer file.Close()

	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"invalidEaterId": {fileName: file.Name(), wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.IncorrectInputErr},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			gotSortedMenuIDs, gotOrderedItemCounts, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
			if !reflect.DeepEqual(gotSortedMenuIDs, tc.wantSortedMenuIDs) {
				t.Errorf("FetchTopThreeOrderedItems() gotSortedMenuIDs = %v, want %v", gotSortedMenuIDs, tc.wantSortedMenuIDs)
			}
			if !reflect.DeepEqual(gotOrderedItemCounts, tc.wantOrderedItemCounts) {
				t.Errorf("FetchTopThreeOrderedItems() gotOrderedItemCounts = %v, want %v", gotOrderedItemCounts, tc.wantOrderedItemCounts)
			}
		})
	}
}

func TestFetchTopThreeOrderedItems_InvalidFoofMenuId(t *testing.T) {

	// Open the log file
	file, err := os.Open("log_test.txt")
	if err != nil {
		err = errors.NewNotFoundError(order.FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	os.WriteFile(file.Name(), []byte("1, 1\n2, A\n"), 0664)
	defer file.Close()

	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"invalidFoofMenuId": {fileName: file.Name(), wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.IncorrectInputErr},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			_, _, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
		})
	}
}

func TestFetchTopThreeOrderedItems_DuplicateEntry(t *testing.T) {

	// Open the log file
	file, err := os.Open("log_test.txt")
	if err != nil {
		err = errors.NewNotFoundError(order.FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	os.WriteFile(file.Name(), []byte("1,1\n2, 2\n1, 1\n"), 0664)
	defer file.Close()

	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"duplicateEntry": {fileName: file.Name(), wantSortedMenuIDs: nil, wantOrderedItemCounts: nil, wantErr: order.IncorrectInputErr},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			_, _, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
		})
	}
}

func TestFetchTopThreeOrderedItems_Success(t *testing.T) {

	// Open the log file
	file, err := os.Open("log_test.txt")
	if err != nil {
		err = errors.NewNotFoundError(order.FileNotFoundErr, fmt.Errorf("File Not Found Error"))
		return
	}
	os.WriteFile(file.Name(), []byte("1, 1\n2, 2\n3, 3\n4, 1\n5, 1\n6, 2\n7, 2\n8, 3\n9, 2\n10, 2\n"), 0664)
	defer file.Close()

	var sortedMenuIDs = []int{2, 1, 3}
	var orderedItemCounts = map[int]int{1:3, 2:5, 3:2}
	tests := map[string]struct {
		fileName              string
		wantSortedMenuIDs     []int
		wantOrderedItemCounts map[int]int
		wantErr               errors.Code
	}{
		"success": {fileName: file.Name(), wantSortedMenuIDs: sortedMenuIDs, wantOrderedItemCounts: orderedItemCounts, wantErr: 0},
	}
	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			gotSortedMenuIDs, gotOrderedItemCounts, err := order.FetchTopThreeOrderedItems(tc.fileName)
			if err != nil {
				error, ok := err.(errors.Error)
				assert.True(t, ok, "")
				assert.Equal(t, tc.wantErr, error.Code())
			}
			if !reflect.DeepEqual(gotSortedMenuIDs, tc.wantSortedMenuIDs) {
				t.Errorf("FetchTopThreeOrderedItems() gotSortedMenuIDs = %v, want %v", gotSortedMenuIDs, tc.wantSortedMenuIDs)
			}
			if !reflect.DeepEqual(gotOrderedItemCounts, tc.wantOrderedItemCounts) {
				t.Errorf("FetchTopThreeOrderedItems() gotOrderedItemCounts = %v, want %v", gotOrderedItemCounts, tc.wantOrderedItemCounts)
			}
		})
	}
}
