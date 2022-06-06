package main
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 333, 34343434

	p := &i

	*p = 33333333

	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(j)

	v := Vertex{1,3333}
	fmt.Println( v.X )

	var a [10]int
	fmt.Println(a)

	var b = a[1:4]
	fmt.Println(b)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	range names

	c := names[0:2]
	d := names[1:4]
	fmt.Println(c, d)
}