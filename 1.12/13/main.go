package main

import "fmt"

/*
Напишите программу, принимающая на вход число N(N≥4), а затем
N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/

func main() {
	var n, v int

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	//s := make([]int, n)

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&v)
		if err != nil {
			return
		}

		if i == 3 {
			fmt.Println(v)
			return
		}
	}
}
