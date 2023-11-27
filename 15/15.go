package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100] //первоначальная строка занимает большой объём памяти, даже после создания среза в первые 100 символов
}

func main() {
  someFunc()
}
*/

var justString string

func createHugeString(size int) string {
	return strings.Repeat("あ", size) // создаем строку из символа あ размера size
}

func someFunc() {
	var builder strings.Builder // использование strings.Builder для быстрой работы со строками тк строки представляют срез байт
	// также позволяет динамически строить строки без лишнего копирования

	fmt.Println(utf8.RuneLen('あ')) //узнать размер символа в байтах

	v := createHugeString(1 << 10) // создаем большую строку размером 1024

	runes := []rune(v) //перевод строки в руны

	builder.WriteString(string(runes[:100])) // ограничиваемся первыми 100 рунами и добавляем к буферу

	justString = builder.String() // присваиваем строку глобальной переменной
}

func main() {
	someFunc()
	fmt.Println(justString)
}
