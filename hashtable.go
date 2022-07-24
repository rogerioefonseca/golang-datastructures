package main

import "fmt"

type person struct {
	name string
	age  int
}

func testHashTable() {
	people := make(map[string]*person)
	person1["1"] := person{name: "Rogerio", age: 38}
	fmt.Printf(person1["1"])
}
func main() {
	fmt.Println("vim-go")
	testHashTable()
}
