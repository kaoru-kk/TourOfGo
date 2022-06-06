package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(" go run os ")

	switch os := runtime.GOOS; os {
		case "dawin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
		}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("gM!")
	case t.Hour() < 18:
		fmt.Println("gA!")
	default:
		fmt.Println("gE!")
	}

}