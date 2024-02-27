package basic

import "fmt"

type Cat struct {
}

func (item Cat) Sound() {
	fmt.Println("meow..")
}

func (item Cat) NumberOfLegs(){
	fmt.Println(4)
}