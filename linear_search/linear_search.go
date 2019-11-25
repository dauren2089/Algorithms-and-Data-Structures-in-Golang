package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(lSearch(list, 5))
}

func lSearch(list []int, n int) bool {

	for i := 0; i < len(list); i ++ {
		if list[i] == n {
			return true
		}
	}
	return false
}
