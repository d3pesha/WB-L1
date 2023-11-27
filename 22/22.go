package main

import (
	"fmt"
	"math"
	"math/big"
)

/*Разработать программу, которая перемножает,
делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.*/

func main() {
	a := big.NewInt(int64(math.Pow(2, 26)))
	b := big.NewInt(int64(math.Pow(2, 20)))

	fmt.Printf("1st number: %d\n2nd number: %d\n\n", a, b)

	res := big.NewInt(0)
	fmt.Println("multiplex:", res.Mul(a, b))
	fmt.Println("division:", res.Div(a, b))
	fmt.Println("sum:", res.Add(a, b))
	fmt.Println("subtract:", res.Sub(a, b))
}
