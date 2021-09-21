package main

func main() {
	a := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
		close(a)
	}()

	for r := range a {
		println(r)
	}

}
