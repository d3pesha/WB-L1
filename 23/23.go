package main

import "fmt"

/*Удалить i-ый элемент из слайса.*/

func removeByIndex(slice []string, index int) []string {

	// создаем новый срез с длиной на 1 меньше
	result := make([]string, len(slice)-1)

	// обрезаем слайс если индекс равен крайним элементам
	if index == 0 {
		result = append([]string{}, slice[1:]...)
		//result = slice[1:]
		return result
	} else if index == len(slice)-1 {
		result = append([]string{}, slice[:index]...)
		return result
	}

	//добавляем элементы до и после индекса элемента
	result = append(slice[:index], slice[index+1:]...)
	return result
}

func removeByIndexCopy(slice []string, index int) []string {
	// создаем новый срез с длиной на 1 меньше
	result := make([]string, len(slice)-1)

	// копируем элементы до удаляемого
	copy(result[:index], slice[:index])

	// копируем элементы после удаляемого
	copy(result[index:], slice[index+1:])

	return result
}

func main() {
	arr := []string{"0", "1", "2", "3", "44", "55", "666"}

	res := removeByIndexCopy(arr, 5)
	res2 := removeByIndex(arr, 2)

	fmt.Println(res)
	fmt.Println(res2)

}
