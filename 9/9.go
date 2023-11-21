package main

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout*/

func first(ch chan int, array []int) {
	for _, i := range array {
		ch <- i
	}
	close(ch)
}

func second(ch chan int) ([]int, []int) {
	var (
		initial []int
		result  []int
	)

	for i := range ch {
		initial = append(initial, i)
		result = append(result, i*2)
	}
	return initial, result
}

/*func main() {
	array := []int{1, 2, 3, 4, 5, 8, 10, 24}
	ch := make(chan int, len(array))
	go first(ch, array)
	initial, result := second(ch)

	fmt.Println("Initial: ", initial)
	fmt.Println("Result: ", result)
}*/
