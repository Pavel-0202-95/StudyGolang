package main

import (
	"fmt"
)

type DifficultData struct {
	a   int
	str string
}

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil { // в recover можно засунуть что угодно любую структуру.
			fmt.Printf("Our struct: %v\n", c)
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	diffStruct := DifficultData{
		a:   10,
		str: "test",
	}
	panic(diffStruct) // прога крашиться ->
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")

	fmt.Println("main() ended2!")
}
