package main

import "fmt"

// Creating struct person with what fields are expected.
type person struct {
	name   string
	age    int
	height float64
	alive  bool
}

// Create a function named createPerson with return type person
func createPerson(name string) person {
	newPerson := person{
		name:   name,
		age:    10,
		height: 0.2,
		alive:  true,
	}
	return newPerson
}

func main() {
	// Create new person and modify the value later on
	a := createPerson("Jerry")
	fmt.Println(a)
	a.age = 1
	fmt.Println(a)
}
