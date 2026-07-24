package main

import "fmt"

/*
Поменяйте местами значения переменных на которые ссылаются указатели.
После этого переменные нужно вывести.
*/

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	test(&x1, &x2)
}
