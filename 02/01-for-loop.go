package main

import "fmt"

func main() {
	// Method 1
	i := 0
	// While True, run Forever
	for {
		if i <= 10 {
			fmt.Println(i)
		} else {
			break
		}
		i++
	}

	// Method 2
	for a := 0; a <= 10; a++ {
		fmt.Println("a-> ", a)
	}

	// Method 3
	b := 0
	for b <= 10 {
		fmt.Println("b-> ", b)
		b++
	}

	// Print even,odd upto 20
	c := 0
	for {
		if c%2 == 0 {
			fmt.Println("Even-> ", c)
			c++
			// continue skips any further loop code and continues next iteration
			continue
		}
		fmt.Println("Odd-> ", c)
		c++

		if c > 20 {
			break
		}
	}

}
