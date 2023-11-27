package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

func reverseWords(str string) string {
	//создаём срез из слов разделенных пробелом
	splitter := strings.Split(str, " ")

	// проходим половину среза
	for i := 0; i < len(splitter)/2; i++ {
		//определяем индекс слова, на противоположной стороне от текущего
		j := len(splitter) - i - 1
		//меняем местами текущее слово на противоположной стороне
		splitter[i], splitter[j] = splitter[j], splitter[i]
	}
	//объединяем перевернутые слова в строку, разделяя пробелами.
	return strings.Join(splitter, " ")
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("hell he a ven world"))
}
