package main

import "fmt"

func main() {
	defer fmt.Println("defer!!")

	fmt.Println("heeeeee")
	
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	defer fmt.Println("last defer!!")
}