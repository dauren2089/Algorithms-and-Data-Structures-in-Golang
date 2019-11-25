package main

import (
	"fmt"
	"sort"
)

func IsThatNumber(list []int, i int) bool {
	low := 0
	high := len(list)-1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return true
		} else if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4, 7}
	fmt.Println(s)
	sort.Ints(s)

	fmt.Println(IsThatNumber(s, 9))
	fmt.Println(IsThatNumber(s, 7))
	fmt.Println(s)

}