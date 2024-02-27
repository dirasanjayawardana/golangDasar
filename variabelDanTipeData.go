package main

import "fmt"

func main() {
	
	// variabel yang sudah dibuat harus digunakan, jika tidak maka akan error
	// format membuat variabel --> var namaVariabel tipeData = value
	// var contohInterger int

	// jika deklarasi variabel langsung diisi nilainya, maka tidak perlu menyebutkan tipe datanya
	// var contohInteger2 = 10
	// contohInteger3 := 10

	// function untuk string, len("string") --> untuk menghitung jumlah karakter distring, "string"[index] --> untuk mengambil karakter pada index tertentu
	// var contohString string = "dira sanjaya"

	// var contohBool bool = true

	// multiple variabel --> membuat beberapa variabel sekaligus
	// var (
	// 	firstName = "Dira"
	// 	lastName  = "Sanjaya"
	// )
	// fmt.Println("hasil print : ", firstName)
	// fmt.Println("hasil print : ", lastName)

	// const --> variabel yang tidak bisa diubah lagi isinya, wajib langsung diisi ketika variabel dibuat, tetapi tidak error meski tidak digunakan
	const contohConstant = "dira sanjaya"
	fmt.Println(contohConstant)
}
