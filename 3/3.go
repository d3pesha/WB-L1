package main

import (
	"fmt"
	"sync"
)

/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов с использованием конкурентных вычислений. */

func square(n int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sq := 1
	for i := 1; i <= n; i++ {
		sq = i * i
	}
	c <- sq
}

func main() {
	var wg sync.WaitGroup
	array := []int{2, 4, 6, 8, 10}
	c := make(chan int, len(array))

	for _, i := range array {
		wg.Add(1)
		go square(i, c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	sum := 0
	for res := range c {
		sum += res
		fmt.Printf("test %d ", sum)
	}

}
