package main

import "fmt"

func work(a int) int {
	return a * 10
}

func main() {

	MyMap := make(map[int]int)

	var a int

	for i := 0; i < 1; i++ {
		fmt.Scanln(&a)

		if _, ok := MyMap[a]; !ok { // ok = isExist
			MyMap[a] = work(a)
		}
		fmt.Print(MyMap[a], " ")
	}

}
