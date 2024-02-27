package basic

import "fmt"

func Main_function() {
	fmt.Println(sum(10, 20))
	sayHello("dira sanjaya")
}

// function dengan return
func sum(number1, number2 int) int {
	return number1 + number2
}

// function tanpa return
func sayHello(name string){
	fmt.Printf("hello %s\n", name)
}