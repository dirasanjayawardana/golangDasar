package basic

import "fmt"

type Dog struct {
}

func (item Dog) Sound() {
	fmt.Println("gug gug..")
}

func (item Dog) NumberOfLegs(){
	fmt.Println(4)
}