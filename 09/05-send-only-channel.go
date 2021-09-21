// Send only channel restricts data receiving only. Useful while restrictions.
package main

import "fmt"

func main() {
	data := make(chan string, 1)

	ping(data, "Sending String")

	fmt.Println(<-data)

}

// send accept channel with type string
func ping(send chan<- string, s string) {
	send <- s
}
