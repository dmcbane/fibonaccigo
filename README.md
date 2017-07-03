# fibonaccigo
Playing around with Fibonacci implementations as I learn the Go language. See the Readme for descriptions of the various implementations.   This is a reimplementation of the Rust version for comparison

When I start to learn a new language, [Go](https://golang.org/) in this case, I usually use small well defined projects that I understand well and attempt to implement them in the new language.  At least that's how it starts.  As I explore the language, it's paradigms and typical approaches to solving problems, I tend to wander from the original concept.  This project is a direct translation of my fibonacci project implemented in [Rust](https://www.rust-lang.org) so that I can compare the two languages as I learn about them.

1. This version of the Fibonacci function uses internal floats and ints to implement [Binet's formula](https://artofproblemsolving.com/wiki/index.php?title=Binet%27s_Formula).  It is extremely fast but is only accurate up to the 75th Fibonacci number.
2. This version uses a brute force iterative implementation with [math.big](https://golang.org/pkg/math/big/#). The iterative algorithm is faster than the fast doubling algorithm for lower numbers, but is quickly surpassed.
3. This version of the Fibonacci function uses [math.big](https://golang.org/pkg/math/big/#) to implement [Binet's formula](https://artofproblemsolving.com/wiki/index.php?title=Binet%27s_Formula).  It is slower than 1 above, but extends the accuracy up to the 281st Fibonacci number.
4. This version (fast doubling) performs much better for large Fibonacci numbers and grows much more slowly than the iterative algorithm.  See https://www.nayuki.io/page/fast-fibonacci-algorithms
