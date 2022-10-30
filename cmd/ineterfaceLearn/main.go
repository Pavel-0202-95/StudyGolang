package main

import "fmt"

type ISwiming interface {
	Swiming()
	WhoIs() string
}

type Duck struct {
	Name   string
	Number int
}

func (d *Duck) Crya() {
	fmt.Println("crya crya")
}

func (d *Duck) Swiming() {
	fmt.Println("Duck swiming")
}

func (d *Duck) WhoIs() string {
	return "Duck"
}

type Swallow struct {
	TimeFly  int
	CountFly int
}

func (s *Swallow) Fly() {
	fmt.Println("Fly")
}

func (s *Swallow) WhoIs() string {
	return "Swallow"
}

func (d *Swallow) Swiming() {
	fmt.Println("Swallow swiming")
}

func WhoSwiming(obj ISwiming) {
	obj.Swiming()
}

func main() {
	duck := &Duck{}
	testssss := map[string]ISwiming{
		"test": duck,
	}
	//WhoSwiming(duck)
	swallow := &Swallow{}
	testssss["dfdsf"] = swallow
	//WhoSwiming(swallow)
	for _, val := range testssss {
		val.Swiming()
		fmt.Println(val.WhoIs())
		if val.WhoIs() == "Duck" {
			du := val.(*Duck)
			du.Crya()
		} else if val.WhoIs() == "Swallow" {
			v := val.(*Swallow)
			v.Fly()
		}
		//if v, ok := val.(*Swallow); ok {
		//	v.Fly()
		//} else {
		//	du := val.(*Duck)
		//	du.Crya()
		//}
		//fmt.Println(key, " ", val)
	}
}
