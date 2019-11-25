package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var list []int
	for i := 0; i <= 100; i ++ {
		list = append(list,rand.Intn(1000))
	}
	fmt.Printf("Unsorted List: %v \n", list)

	SelectSort(list)
	InsertionSort(list)
	ShellSort(list)
	BubbleSort2(list)
	BubbleSort3(list)
	fmt.Printf("Quick sort: %v \n", Qwik_Sort(list,0,len(list)-1))
}

func SelectSort(list []int) {
	for i := 0; i <= len(list); i++ {
		for j := i + 1; j < len(list); j ++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
				//fmt.Printf("Iteration: %v, %v", i, list)
			}
		}
	}
	fmt.Printf("Select Sort: %v \n", list)
}

func InsertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		x := list[i]
		j := i
		for ; j >= 1 && list[j-1] > x; j-- {
			list[j] = list[j-1]
		}
		list[j] = x
	}

	fmt.Printf("Insertion Sort: %v \n", list)
}

func ShellSort(list []int) {
	for gap := len(list) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(list); i++ {
			x := list[i]
			j := i
			for ; j >= gap && list[j-gap] > x; j -= gap {
				list[j] = list[j-gap]
			}
			list[j] = x
		}
	}
	fmt.Printf("Shell Sort: %v \n", list)
}

func BubbleSort2(list []int) {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	fmt.Printf("Bubble Sort v2: %v \n", list)
}

func BubbleSort3(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	fmt.Printf("Bubble Sort v3: %v \n", list)
}

func Qwik_Sort(lst []int,left int,right int) []int {

	//Создаем копии пришедших переменных, с которыми будем манипулировать в дальнейшем.
	l := left
	r := right

	//Вычисляем 'центр', на который будем опираться. Берем значение ~центральной ячейки массива.
	center := lst[(left + right) / 2]

	//fmt.Println(l,r,(left + right) / 2)

	//Цикл, начинающий саму сортировку
	for l <= r{

		//Ищем значения больше 'центра'
		for lst[r] > center{
			r--
		}

		//Ищем значения меньше 'центра'
		for lst[l] < center{
			l++
		}

		//После прохода циклов проверяем счетчики циклов
		if(l <= r){
			//И если условие true, то меняем ячейки друг с другом.
			lst[r],lst[l] = lst[l],lst[r]
			l++
			r--
		}
	}

	if r > left{
		Qwik_Sort(lst,left,r)
	}

	if l<right{
		Qwik_Sort(lst,l,right)
	}

	return lst
}

