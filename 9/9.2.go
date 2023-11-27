package main

import (
	"context"
	"fmt"
	"time"
)

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/

func main() {

	var (
		ch1 = make(chan int)
		ch2 = make(chan int)
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case ch2 <- (<-ch1 * 2):
			case <-ctx.Done():
				return

			}
		}
	}()

	go func() {
		for {
			select {
			case x := <-ch2:
				fmt.Println(x)

			case <-ctx.Done():
				return
			}

		}
	}()

	for i := 0; ; i++ {
		ch1 <- i
		time.Sleep(time.Second)
	}
}
