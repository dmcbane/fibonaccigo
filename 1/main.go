package main

import (
	"fmt"
	"github.com/dmcbane/fib/fibpkg"
)

func main() {
	for i := uint32(0); i < uint32(76); i++ {
		fmt.Printf("fib(%2v) = %16.16v\n", i, fibpkg.FibBinet(i))
	}
}
