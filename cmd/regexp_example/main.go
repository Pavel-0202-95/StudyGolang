package main

import (
	"fmt"
	"regexp"
)

func magicFunction(str string) bool {
	m, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)
	return m
}

func main() {
	var text string = "fdsghdfgjsdDD1"
	//fmt.Scan(&text)

	if len(text) >= 5 && magicFunction(text) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}
