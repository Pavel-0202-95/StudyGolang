package main

import "fmt"

func main() {
	condition := true
	param := 1
	//for ok := true; ok; ok = condition {
	//	param++
	//	if param >= 10 {
	//		ok = false
	//	}
	//	fmt.Print(param, " ")
	//}

	//for i := 0; i < 5; i++ {
	//	fmt.Print(i, " ")
	//}

	for ok := true; ok; ok = (param <= 10) {
		param++

		//if param >= 10 {
		//	//param++
		//	ok = false
		//} else {
		//	fmt.Print(param, " ")
		//}
		fmt.Print(param, " ")
	}

	param1 := 0

	for ok := true; ok; ok = condition { // ok = condition - выполняется после итерации
		param1++

		if param1 >= 10 {
			condition = false
		}

		fmt.Print(param1, " ")
	}

}
