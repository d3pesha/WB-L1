package main

import "fmt"

/*Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout. */

func square(n int, c chan int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = i * i
	}
	c <- res

}

func main() {
	array := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	for _, i := range array {
		go square(i, c)
	}
	for range array {
		fmt.Printf("test %d ", <-c)
	}
}
