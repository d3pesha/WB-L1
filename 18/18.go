package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) increment() { // функция инкрементации
	defer c.mutex.Unlock()
	c.value++
	c.mutex.Lock()

}

func main() {

	counter := Counter{value: 0} // экземпляр структуры

	//объявление вэйтгруппы для синхронизации
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	// обработка сигнала прерывания для отмены контекста
	go func() {
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, os.Interrupt)
		<-quitChan
		cancel() // Отмена контекста при получении сигнала прерывания
	}()

	//создание 10 воркеров
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return // Выход из горутины при завершении контекста
				default:
					counter.increment() // Вызов функции инкрементации
				}
			}
		}()
	}

	wg.Wait() // Ожидание завершения всех воркеров

	fmt.Println(counter.value) // Вывод итогового значения счетчика
}
