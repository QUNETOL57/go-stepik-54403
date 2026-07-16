package main

import "fmt"

/*
Найдите самое большее число на отрезке от a до b, кратное 7.

---Входные данные
Вводится два целых числа a и b (a≤b).

---Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b),
кратное 7, или выведите "NO" - если таковых нет.
*/

func main() {
	var a, b, r int
	_, err := fmt.Scan(&a, &b)
	if err != nil || a > b {
		return
	}
	for i := b; i >= a; i -= 1 {
		if i%7 == 0 {
			r = i
			fmt.Println(r)
			return
		}
	}
	fmt.Println("NO")
}
