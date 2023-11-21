package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться. */

func sendData(ch chan<- int) { // создание функции записи в канал

	defer close(ch)     // закрыть канал при завершении записи
	for i := 0; ; i++ { // бесконечный цикл записи
		ch <- i                     //запись в канал
		time.Sleep(1 * time.Second) //секундная пауза для удобного отслеживания таймера
	}
}

func readData(ch <-chan int) { // функция чтения из канала
	for {
		data, ok := <-ch //бесконечный цикл чтения из канала
		if !ok {         // если канал закрыт, выход из цикла
			break
		}
		log.Println("Received:", data) // вывод данных с отслеживанием времени
	}
}

func main() {
	ch := make(chan int)
	var sec int
	fmt.Println("Enter time long in second:")
	_, err := fmt.Fscan(os.Stdin, &sec) //ввод продолжительности работы программы
	if err != nil {
		log.Println(err) //обработка ошибки ввода
		return
	}
	go sendData(ch)
	go readData(ch)                              //запуск двух горутин на запись и чтение
	time.Sleep(time.Duration(sec) * time.Second) //приостановление программы через sec
}
