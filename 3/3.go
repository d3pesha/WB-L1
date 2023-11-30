package main

import (
	"fmt"
	"sync"
)

/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов с использованием конкурентных вычислений. */

func square(n int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика вг

	res := n * n // Вычисление квадрата числа
	c <- res     // Отправка результата в канал
}

func main() {
	var wg sync.WaitGroup
	array := []int{2, 4, 6, 8, 10}
	c := make(chan int, len(array))

	// Запуск горутин для вычисления квадрата чисел в массиве
	for _, i := range array {
		wg.Add(1) // Увеличение счетчика вг перед запуском горутины
		go square(i, c, &wg)
	}

	// Запуск горутины для ожидания завершения всех горутин
	go func() {
		wg.Wait()
		close(c) //закрытие канала
	}()

	// Суммирование результатов из канала и вывод их в консоль
	sum := 0
	for res := range c {
		sum += res
		fmt.Printf("%d ", sum)
	}

}
