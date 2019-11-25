package main_test

import (
	"fmt"
	"testing"
)

func TestisThatNumber(list []int, i int) bool{
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return true
			break
		} else if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 45, 34, 65, 78, 98, 21, 23, 54, 55, 12, 35, 65, 88, 98, 100, 21, 45, 34, 33, 101}

	res := TestisThatNumber(list, 77)
	expected := true
	if res == expected {
		t.Errorf("Our binary search function doens't work, %v isn't %v\n", res, expected)
	}
	fmt.Println(res)
}