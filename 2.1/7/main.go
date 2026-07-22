package main

import "fmt"

/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

---Входные данные
Вводится четыре числа.

---Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел
*/

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	switch {
	case a < b && a < c && a < d:
		return a
	case b < a && b < c && b < d:
		return b
	case c < a && c < b && c < d:
		return c
	default:
		return d
	}
}

func main() {
	r := minimumFromFour()
	fmt.Println(r)
}
