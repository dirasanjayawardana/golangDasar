package basic

import "fmt"

// * --> untuk mengambil nilai dari sebuah address memory
// & --> menghasilkan address memory dari sebuah variabel

func Pointer() {
	// variable pointer --> mereferensikan suatu variabel ke address memory variabel lain
	var tanggal int = 20           // variabel biasa
	var pointerTgl *int = &tanggal // variabel pointer

	fmt.Println(tanggal, *pointerTgl)
	fmt.Println(&tanggal, pointerTgl)
}