package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	NickNames []string
}

func main() {
	b := []byte(`{"name": "John", "age": 30, "nicknames": ["John", "Johny", "Johnathan"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.NickNames)
}
