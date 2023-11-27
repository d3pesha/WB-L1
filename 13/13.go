package main

import "fmt"

/*Поменять местами два числа без создания временной переменной.*/

func main() {
	x := 5
	y := 10

	fmt.Println(swapWithArithmetic(x, y))
	fmt.Println(swapWithXOR(x, y))
	fmt.Println(swapWithAssigment(x, y))
}

func swapWithArithmetic(x, y int) (int, int) {
	// Выполняем операции сложения и вычитания для обмена значениями между переменными.
	x = x + y
	y = x - y
	x = x - y

	// Возвращаем измененные значения переменных.
	return x, y
}

func swapWithXOR(x, y int) (int, int) {
	//бинарное представление 5 = 0101, 10 = 1010

	x = x ^ y // теперь х = 1111
	y = x ^ y // у = 0101
	x = x ^ y // х = 1010

	return x, y
}

func swapWithAssigment(x, y int) (int, int) {
	x, y = y, x //использование множественного присваивания
	return x, y
}
