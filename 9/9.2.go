package main

import (
	"context"
	"fmt"
	"time"
)

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/

func main() {

	// объявление двух каналов
	var (
		ch1 = make(chan int)
		ch2 = make(chan int)
	)

	// добавление контекста для выхода
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Горутина для умножения чисел на 2 и записи во второй канал
	go func() {
		defer close(ch2)
		for {
			select {
			default:
				ch2 <- (<-ch1 * 2) //запись в канал данных
			case <-ctx.Done(): //выход из горутины
				return

			}

		}
	}()
	// чтение из второго канала и вывода данных
	go func() {
		for {
			select {
			case x, ok := <-ch2:
				if !ok {
					return // если канал закрыт - выход
				}
				fmt.Println(x) // вывод

			case <-ctx.Done():
				return // Выход из горутины при получении сигнала отмены
			}

		}
	}()

	// запись чисел в первый канал
	for i := 0; ; i++ {
		ch1 <- i
		time.Sleep(time.Second)
	}
}
