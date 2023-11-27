package main

import (
	"fmt"
	"strings"
)

// целевой интерфейс, с которым работает клиент
type Target interface {
	GetStr(len int) string
}

// адаптируемый интерефейс, который используется через адаптер
type Adaptee interface {
	GetRune(len int) []rune
}

// HugeStr реализует интерфейс Target
type HugeStr struct{}

// создание строки
func (h *HugeStr) GetStr(len int) string {
	return strings.Repeat("q", len)
}

// адаптация к интерфейсу Adaptee
type Adapter struct {
	target Target
}

// реализация метода интерфейса Adaptee
func (a *Adapter) GetRune(len int) []rune {
	return []rune(a.target.GetStr(len))
}

func main() {
	// создание экземпляра адаптера, который адаптирует HugeStr к интерфейсу Adaptee
	var adaptee Adaptee = &Adapter{target: &HugeStr{}}

	// использование адаптера для получения среза рун и изменение одной из них
	temp := adaptee.GetRune(10)
	temp[2] = 'x'
	
	fmt.Println(string(temp))
}
