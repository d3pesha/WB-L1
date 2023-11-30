package main

import (
	"fmt"
	"sync"
)

/*Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout. */

// функция, вычисляющая квадрат числа и отправляющая результат в канал
func square(n int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := n * n
	c <- res
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	var wg sync.WaitGroup //для ожидания завершения всех горутин

	// запуск горутин для вычисления квадратов чисел
	for _, i := range array {
		wg.Add(1) // увеличение счетчика WaitGroup перед запуском горутины
		go square(i, c, &wg)
	}

	// горутина, которая ждёт заверешения всех горутин и закрывает канал
	go func() {
		wg.Wait()
		close(c)
	}()
	// чтение результатов из канала и вывод
	for result := range c {
		fmt.Printf("%d ", result)
	}
}
