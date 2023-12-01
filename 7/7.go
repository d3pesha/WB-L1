package main

import (
	"fmt"
	"sync"
)

/*Реализовать конкурентную запись данных в map.*/

type maps struct {
	mute sync.Mutex // мьютекс для синхронизации доступа к мапе
	m    map[string]int
}

// создание нового экземпляра мапы
func newMap() *maps {
	return &maps{
		m: make(map[string]int),
	}
}

// установка значения по ключу
func (mp *maps) set(key string, value int) {
	mp.mute.Lock()
	defer mp.mute.Unlock()
	mp.m[key] = value
}

// получение значения по ключу из мапы
func (mp *maps) get(key string) int {
	mp.mute.Lock()
	defer mp.mute.Unlock()
	return mp.m[key]
}

func main() {
	m := newMap() // новый экземпляр
	var wg sync.WaitGroup

	// запуск 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)            //добавление счетчика вг
		go func(index int) { //использование замыкания с параметром index
			defer wg.Done() //уменьшение счетчика вг
			key := fmt.Sprintf("k%d", index)
			value := index * 2
			m.set(key, value) // установка ключ-значения в мапу
			fmt.Printf("Goroutine %d set [%s]%d\n", index, key, value)
		}(i)
	}
	wg.Wait() // ждём завершения горутин

	// вывод данных из мапы
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("k%d", i)
		value := m.get(key)
		fmt.Printf("Value at %s: %d\n", key, value)
	}
}
