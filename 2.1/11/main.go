package main

import "fmt"

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

func sumInt(v ...int) (int, int) {
	var sum, count int
	for _, v := range v {
		sum += v
		count++
	}
	return count, sum
}

func main() {
	count, sum := sumInt(1, 2, 4)
	fmt.Println(count, sum)
}
