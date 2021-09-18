package main

import "fmt"

// Creating Interface
type personData interface {
	// Define method signatures provided by interface
	speakHello() string
	greet() string
}

type person struct {
	name   string
	age    int
	height float64
	alive  bool
}

func (p person) speakHello() string {
	return "Hi I'm " + p.name
}

func (p person) greet() string {
	return "Greetings from " + p.name
}

func main() {
	var person1 personData
	person1 = person{
		name:   "Kunal",
		age:    33,
		height: 5.11,
		alive:  true,
	}
	fmt.Println(person1.speakHello())
	fmt.Println(person1.greet())

}
