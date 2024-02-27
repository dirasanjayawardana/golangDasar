package basic

import "fmt"

type User struct {
	firstName string
	lastName  string
	Age       int
}

// constructor
func NewUser(firstName, lastName string, age int) *User {
	return &User{firstName, lastName, age}
}

// method
func (item *User) GetFullName() string {
	return fmt.Sprintf("%s %s", item.firstName, item.lastName)
}