package main

import "fmt"

func main() {
	t, f := true, false

	fmt.Printf("%t %v\n", t, f)
	fmt.Printf("%t %v\n", f, f)
}
