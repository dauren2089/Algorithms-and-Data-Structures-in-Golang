package main_test

import (
	"fmt"
	"testing"
)

func main(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 45, 34, 65, 78, 98, 21, 23, 54, 55, 12, 35, 65, 88, 98, 100, 21, 45, 34, 33, 101}
	res := TestlSearch(list, 77)

	expected := true
	if res != expected {
		t.Errorf("Our linear search function doens't work correctly")
	}
	fmt.Println(res)
}

func TestlSearch(list []int, n int) bool {

	for i := 0; i < len(list)-1; i ++ {
		if list[i] == n {
			return true
		}
	}
	return false
}