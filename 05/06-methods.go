package main

import "fmt"

type dimensions struct {
	length, width int
}

func (r dimensions) areaRect() int {
	return r.width * r.length
}

func (r dimensions) perimRect() int {
	return 2*r.width + 2*r.length
}

func main() {
	r := dimensions{10, 2}

	fmt.Println(r.areaRect())
	fmt.Println(r.perimRect())

}
