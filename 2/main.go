package main

import (
	"fmt"
	"github.com/dmcbane/fib/fibpkg"
)

func main() {
	for i := uint32(0); i < uint32(10001); i++ {
		fmt.Printf("fib(%5v) = %v\n", i, fibpkg.FibIterativeBigint(i))
	}
}
