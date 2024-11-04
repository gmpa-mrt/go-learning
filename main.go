package main

import (
	"fmt"
	"github.com/gmpa-mrt/go-learning/Fibonacci"
)

func main() {
	f := Fibonacci.Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
