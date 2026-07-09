package main

import "fmt"

/*
Заданы три числа -a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil || a > b || b > c {
		return
	}
	if c*c == a*a+b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
