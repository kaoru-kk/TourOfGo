package main

import (
	"fmt"
	"math"
)


func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum -= i
	// }
	sum := 2
	for sum < 100 {
		sum += sum

		if (sum % 32) == 0 {
			fmt.Println("32!!")
		} else {
			fmt.Println("no")
		}
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 2),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

