package main

import (
	"fmt"
)

/*Реализовать пересечение двух неупорядоченных множеств.*/

// использование дженериков
func insert[T comparable](set1, set2 []T) []T {
	var result []T // объявление слайса

	check := make(map[T]bool) //объявление мапы для проверки существования элемента

	//задать каждому элементу из set1 значение true
	for _, i := range set1 {
		check[i] = true
	}

	//пройти по каждому элементу set2
	for _, v := range set2 {
		if check[v] { // если элемент существует в set1
			result = append(result, v) //добавить в результат
		}
	}
	return result

}

func main() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 6, 7, 2, 8}

	fmt.Println(insert(s1, s2))

	s3 := []string{"1", "2", "j", "s", "ka"}
	s4 := []string{"ka", "2", "1"}

	fmt.Println(insert(s3, s4))
}
