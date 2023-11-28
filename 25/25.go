package main

import (
	"log"
	"time"
)

/*Реализовать собственную функцию sleep*/

// используем метод After, в котором через определенное время поступает сигнал из канала
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	sec := 3
	log.Println("Начало работы")
	sleep(sec)
	log.Printf("Завершение работы после ожидания %d секунд", sec)

}
