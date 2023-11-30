package main

import "fmt"

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).*/

type Human struct {
	message string
}

type Action struct {
	Human   // встраивание структуры Human
	message string
}

func (h *Human) sendMessage() {
	fmt.Printf("message by human: %s\n", h.message)
}

func (a *Action) sendMessageInAction() {
	fmt.Printf("message by action: %s", a.message)
}

func main() {
	// создание объекта Action с встроенным объектом Human
	action := Action{
		Human:   Human{message: "Hello"}, // инициализация встроенного Human
		message: "world",                 // инициализация message в Action
	}
	// вызов методов
	action.sendMessage()
	action.sendMessageInAction()
}
