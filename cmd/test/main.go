package main

import (
	"fmt"
	"github.com/Pavel-0202-95/StudyGolang/Internal/math"
	uuid "github.com/nu7hatch/gouuid"
	"log"
)

func main(){
	fmt.Println(math.Abs(-10))

	str, err := uuid.NewV4()

	if err!=nil{
		log.Panic(err)
	}

	fmt.Println(str.String())
}
