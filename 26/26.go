package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой. Например:
abcd — true
abCdefAaf — false
aabcd — false */

func unique(str string) bool {
	//перевод строки в нижний регистр
	str = strings.ToLower(str)

	//переводим строку в срез рун для удобного сравнения
	runes := []rune(str)

	//создаём вложенные циклы
	for i := 0; i < len(runes)-1; i++ {

		//сравнение текущего символа со всеми последующими символами
		for j := i + 1; j < len(runes); j++ {
			// если найден дубликат символа, возвращается false
			if runes[i] == runes[j] {
				return false

			}
		}

	}
	//если дубликатов нет - true
	return true
}

func main() {
	fmt.Println(unique("abcd"))      //true
	fmt.Println(unique("abCdefAaf")) //false
	fmt.Println(unique("aabcd"))     //false
	fmt.Println(unique("add"))       //false
	fmt.Println(unique("dd"))        //false
	fmt.Println(unique(""))          //true

}
