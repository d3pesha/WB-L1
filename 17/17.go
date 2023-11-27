package main

import "fmt"

/*Реализовать бинарный поиск встроенными методами языка.*/

func binarySearch(arr []int, key int) int {

	low, high := 0, len(arr)-1 //задать границы

	for low <= high {
		mid := (low + high) / 2 //поиск середины массива

		if arr[mid] == key {
			return mid //элемент найден
		} else if arr[mid] < key {
			low = mid + 1 //искать в правой половине со сдвигом
		} else {
			high = mid - 1 //искать в левой половине со сдвигом
		}
	}
	return -1 //элемент не найден
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 7

	index := binarySearch(numbers, x)

	if index == -1 {
		fmt.Printf("Элемент %d не найден\n", x)
	} else {
		fmt.Printf("Элемент %d найден в позиции %d\n", x, index)
	}
}
