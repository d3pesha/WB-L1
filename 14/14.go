package main

import (
	"fmt"
	"reflect"
)

/*Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.*/

func main() {
	var intValue int = 42
	var stringValue string = "Hello, World!"
	var boolValue bool = true
	var chanValue chan int = make(chan int)

	fmt.Printf("Type of %d: %s\n", intValue, getType(intValue))
	fmt.Printf("Type of %s: %s \n", stringValue, getType(stringValue))
	fmt.Printf("Type of %t: %s \n", boolValue, getType(boolValue))
	fmt.Printf("Type of %v: %s \n", chanValue, getType(chanValue))

}

func getType(v interface{}) string {
	t := reflect.TypeOf(v)
	return t.String()
}
