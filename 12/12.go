package main

import "fmt"

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.*/

func main() {
	example := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем мапу set, которая представляет собой множество строк
	set := make(map[string]bool)

	// проходим по каждому элементу в example
	for _, i := range example {
		// устанавливаем значение true для каждого элемента в мапе set
		set[i] = true
	}

	// проходим по ключам в мапе set
	for k, _ := range set {
		// выводим каждый уникальный элемент множества
		fmt.Printf("Set: %s \n", k)
	}

}
