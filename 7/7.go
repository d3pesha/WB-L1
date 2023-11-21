package main

import (
	"fmt"
	"sync"
)

/*Реализовать конкурентную запись данных в map.*/

type maps struct {
	mute sync.Mutex
	m    map[string]int
}

func newMap() *maps {
	return &maps{
		m: make(map[string]int),
	}
}

func (mp *maps) set(key string, value int) {
	mp.mute.Lock()
	defer mp.mute.Unlock()
	mp.m[key] = value
}

func (mp *maps) get(key string) int {
	mp.mute.Lock()
	defer mp.mute.Unlock()
	return mp.m[key]
}

func main() {
	m := newMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", index)
			value := index * 2
			m.set(key, value)
			fmt.Printf("Goroutine %d set [%s]%d\n", index, key, value)
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := m.get(key)
		fmt.Printf("Value at %s: %d\n", key, value)
	}
}
