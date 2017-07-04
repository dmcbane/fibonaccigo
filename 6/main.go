package main

import (
	"fmt"
	"github.com/dmcbane/fib/fibpkg"
)

func main() {
	gr := fibpkg.GoldenRatioGenerator()
	for i := 0; i < 100; i++ {
		tmp := gr()
		ratlen := len(tmp.RatString())
		fmt.Printf("golden ratio %3v =\n  fraction:%v\n  decimal:%v\n", i, tmp, tmp.FloatString(ratlen))
	}
}
