package main

import (
	"fmt"
	"time"
	"os/user"
)

func main() {
	fmt.Println("Hello, World!", "TEST TEST")
	fmt.Println(time.Now())
	fmt.Println(user.Current())
}
