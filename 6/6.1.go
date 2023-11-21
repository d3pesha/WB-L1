package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*Реализовать все возможные способы остановки выполнения горутины. */

func example1(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker 1 received stop signal. Exiting.")
			return
		default:
			// Работа горутины
			fmt.Println("1 - Working...")
			time.Sleep(time.Second)
		}
	}
}

func example2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker 2 received stop signal. Exiting.")
			return
		default:
			// Работа горутины
			fmt.Println("2 - Working...")
			time.Sleep(time.Second)
		}
	}
}

func example3() {
	fmt.Println("3 - Working...")
	fmt.Println("Worker 3 received stop signal. Exiting.")
	runtime.Goexit()
}

func main() {
	stop1 := make(chan bool)
	stop2 := make(chan struct{})

	fmt.Println("Choose a goroutine to run (1, 2, 3):")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		go example1(stop1)
		time.Sleep(2 * time.Second)
	case 2:
		ctx, cancel := context.WithCancel(context.Background())
		go example2(ctx)
		time.Sleep(2 * time.Second)
		cancel() // Отменить контекст после некоторого времени
	case 3:
		go example3()
	default:
		fmt.Println("Invalid choice.")
		return
	}

	// Отправить сигнал остановки в канал
	switch choice {
	case 1:
		stop1 <- true
	case 2:
		close(stop2)
	}

	// Подождать завершения горутины
	time.Sleep(time.Second)
	fmt.Println("Main goroutine exiting.")
}
