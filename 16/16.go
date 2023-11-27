package main

import "fmt"

/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/

func main() {
	numbers := []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Println(numbers)
	fmt.Println(quickSort(numbers))
}
func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]

	var less, greater []int
	for _, i := range arr {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}
