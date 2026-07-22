package main

import "fmt"

func vote(x int, y int, z int) int {
	var c int
	if x == 0 {
		c++
	}
	if y == 0 {
		c++
	}
	if z == 0 {
		c++
	}
	if c > 1 {
		return 0
	}
	return 1

}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	r := vote(x, y, z)
	fmt.Println(r)
}
