package main

import (
	"fmt"
)

func main() {
	var (
		a int = 1
		b float64 = 1.2
		c string = "test"
		d, e bool = true, false
	)

	fmt.Println(a, b, c, d, e)
	x1 := 1
	fmt.Printf("%T", x1)
}
