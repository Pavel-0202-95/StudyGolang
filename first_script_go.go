package main

import (
	"fmt"
)

/*
func CreateArr(arr[]) {

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}

}
*/

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print("i: ", i, " ")
	}

	fmt.Println("Typed up a Println")

	//arr1 = [4]int{1, 2, 3, 4}

	//CreateArr(arr1)

	fmt.Print("Write Dock file with info")

	fmt.Print("Write Dock file with info")
	fmt.Print("Write Dock file with info")
	fmt.Print("Write Dock file with info")

	// здесь добавили массив и заполнили его, а также вывели.

	var arr1 []int

	for i := 1; i < 11; i++ {
		arr1 = append(arr1, i)
	}

	fmt.Println(arr1, len(arr1)) // Add new method len(arr1)

}
