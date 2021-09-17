package main

import "fmt"

// Creating struct person with what fields are expected.
type person struct {
	name   string
	age    int
	height float64
	alive  bool
}

func main() {
	// Creating p1 as person
	p1 := person{
		name:   "Kunal",
		age:    33,
		height: 5.11,
		alive:  true,
	}
	fmt.Println(p1)
	// Shorthand, but sequence should be proper as defined in struct.
	p2 := person{
		"Test",
		22,
		5.4,
		true,
	}
	fmt.Println(p2)

	// Skip few fields, they will be assigned with default/zero values
	p3 := person{name: "Ram"}
	fmt.Println(p3)
	p4 := person{
		name:  "Hanuman",
		alive: true,
	}
	fmt.Println(p4)

	// We can add/update fields later too
	p3.height = 6.0
	fmt.Println(p3)
}
