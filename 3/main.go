package main

import (
	"fmt"
	"github.com/dmcbane/fib/fibpkg"
)

func main() {
	for i := uint32(0); i < uint32(282); i++ {
		fmt.Printf("fib(%3v) = %v\n", i, fibpkg.FibBinetBigint(i))
	}
}
