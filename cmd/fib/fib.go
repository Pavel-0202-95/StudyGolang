package main

import (
	"fmt"
	"os"
)

func fib(n int) int {

	if n == 1 || n == 2 {
		return 1
	} else if n == 0 {
		return 0
	}
	n1 := n - 1
	n2 := n - 2
	return fib(n1) + fib(n2)
}

func main() {
	fmt.Print(fib(5), " ")

	arg := os.Args

	if len(arg) == 1 {
		fmt.Print("Please provide one sometimes")
		os.Exit(1)
	} else {
		fmt.Print(arg)
	}

}
