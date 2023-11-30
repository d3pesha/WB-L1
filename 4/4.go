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

// функция для генерации случайной строки заданной длины
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

	//Запрос количества воркеров у пользователя
	fmt.Println("Enter count workers:")
	fmt.Fscan(os.Stdin, &numWorkers)

	dataChannel := make(chan interface{})  //канал для передачи данных
	quitChannel := make(chan os.Signal, 1) // канал для получения сигнала завершения
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)

	// поток, записывающий данные в канал
	go func() {
		for {
			select {
			case <-quitChannel:
				close(dataChannel) // закрытие канала данных при получении сигнала завершения
				wg.Wait()          // ожидание завершения всех воркеров перед завершением главного потока
				return
			default:
				dataChannel <- randomString(rand.Intn(7)) // запись случайных символов
				time.Sleep(1 * time.Second)               // периодичность записи
			}
		}
	}()

	// создание воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // увеличение счетчика вг
		go func(workerID int) {
			defer wg.Done() // уменьшение счетчика вг при завершении горутины
			log.Println("gorutine created: ", workerID)

			// бесконечный цикл для чтения данных из канала
			for data := range dataChannel {
				log.Printf("worker %d have data: %v \n", workerID, data)
			}
		}(i) //передача i в качестве параметра workerID горутине
	}

	// ждём завершения горутин
	wg.Wait()
	fmt.Println("Main thread finished")
}
