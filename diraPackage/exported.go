package dirapackage

import "fmt"

// contoh exported function --> bisa diakses di package yang berbeda --> menggunakan huruf besar diawal
func GetFullName(firstName string, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName) // membuat string baru dengan string template
}
