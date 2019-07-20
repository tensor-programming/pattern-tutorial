package main

import "fmt"

func fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}

	}()

	return out
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

func main() {
	for x := range fib(10000000) {
		fmt.Println(x)
	}
}
