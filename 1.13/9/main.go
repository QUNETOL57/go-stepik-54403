package main

import "fmt"

/*
Найдите количество минимальных элементов в последовательности.

---Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

---Выходные данные
Выведите количество минимальных элементов последовательности.
*/

func main() {
	var n, v, m int
	c := 1
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		return
	}

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&v)
		if err != nil {
			return
		}
		if i == 0 {
			m = v
			continue
		}
		if v < m {
			m = v
			c = 1
		} else if v == m {
			c++
		}
	}
	fmt.Println(c)
}
