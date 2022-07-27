package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLefgs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "oioi"
}

func (d *Dog) NumberOfLegs() string {
	return "4"
}

type Gorilla struct {
	Name string
}

func main() {
	fmt.Println("vim-go")

	dog := Dog{Name: "Rogerio", Breed: "4"}

	fmt.Println(dog)
}
