package main

import "fmt"

/*
Из натурального числа удалить заданную цифру.

---Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

---Выходные данные
Вывести число без заданных цифр.
*/

func main() {
	var n, v int
	var r string
	_, err := fmt.Scan(&n, &v)
	if err != nil {
		return
	}
	for {
		if n == 0 {
			break
		} else if n%10 != v {
			r = fmt.Sprintf("%d%s", n%10, r)
		}
		n = n / 10
	}
	fmt.Println(r)
}
