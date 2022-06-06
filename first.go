package main

import (
	"fmt"
	// "math/rand"
)

func add(x, y string) string {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// "naked" return
func split(sum int) (x, y int) {
	x = sum * 3 / 2
	y = sum - x
	return
}

var a, b, c bool

// this is error: non-declaration statement outside function body
// aaaaa := 1111

func main() {
	// fmt.Println("aaaa", rand.Intn( 10) )
	fmt.Println(add("aaaa", "vvvvv"))

	fmt.Println(swap("1aaaa", "2vvvvv"))

	fmt.Println(split(30))

	var intdayo int
	var d, e int = 1, 2
	var f, g, h = true, true, "no"
	i := 55
	fmt.Println(intdayo, a, b, c, d, e, f, g, h, i)
}