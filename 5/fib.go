package main

import "fmt"

func fibonacci(c chan uint64, quit chan int) {
	x, y := uint64(0), uint64(1)
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan uint64)
	quit := make(chan int)
	go func() {
		for i := uint64(0); i < uint64(94); i++ {
			fmt.Printf("%22v\n", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
