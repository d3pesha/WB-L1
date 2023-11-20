package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.*/

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	var wg sync.WaitGroup
	var numWorkers int

	fmt.Println("Enter count workers:")
	fmt.Fscan(os.Stdin, &numWorkers)

	dataChannel := make(chan interface{})
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-quitChannel:
				close(dataChannel)
				wg.Wait()
				return
			default:
				dataChannel <- randomString(rand.Intn(7))
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			log.Println("gorutine created: ", workerID)
			for data := range dataChannel {
				log.Printf("worker %d have data: %v \n", workerID, data)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Main thread finished")
}
