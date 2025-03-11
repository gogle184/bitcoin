package main

import "fmt"

type Human interface {
  Say() string
}

type Person struct {
  Name string
}

func (p *Person) Say() string {
  return "Hello, my name is " + p.Name
}

func main() {
  var h Human = &Person{Name: "John"}
  fmt.Println(h.Say())
}
