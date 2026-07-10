package main

import "fmt"

/*
По данным числам, определите количество чисел, которые равны нулю.

---Входные данные
Вводится натуральное число N, а затем N чисел.

---Выходные данные
Выведите количество чисел, которые равны нулю.
*/

func main() {
	var n, v, c int
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		return
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&v)
		if err != nil {
			return
		}
		if v == 0 {
			c++
		}
	}
	fmt.Println(c)
}
