package main

import (
	basic "dirasanjayawardana/golangDasar/basic"
	dirapackage "dirasanjayawardana/golangDasar/diraPackage"
	"fmt"
)

// go run . --> untuk running semua file.go di dalam direktory
// go run file1.go file2.go --> untuk run beberapa file.go

// unexported --> tidak bisa diakses dari package yang berbeda --> menggunakan huruf kecil diawal
// exported --> bisa diakses dari package yang berbeda --> menggunakan huruf besar diawal
func main() {

	// basic.HelloWorld()
	// basic.VariabelDanTipeData()
	// basic.Main_function()
	// basic.Percabangan()
	// basic.Perulangan()

	// mengakses function lain dari package yang sama dengan unexported function
	fmt.Println(isOdd(10))

	// mengakses exported function dari package yang berbeda
	fmt.Println(dirapackage.GetFullName("dira", "sanjaya"))

	// mengakses variabel pointer
	basic.Pointer()

	// // membuat struct struct (mirip seperti object dan class)
	// var mypost basic.Post
	// mypost.Title = "judul artikel"
	// mypost.Content = "Lorem ipsum"
	// mypost.Owner = "dira sanjaya"
	// fmt.Println(mypost)

	// // cara lain membuat struct (dengan mengisi langsung isi property nya)
	// mypost2 := basic.Post{
	// 	Title: "judul artikel 2",
	// 	Content: "lorem ipsum 2",
	// 	Owner: "dira sanjaya 2",
	// }
	// fmt.Println(mypost2)

	// membuat struct dengan constructor/paramter
	myuser := basic.NewUser("dira", "sanjaya", 24)
	mypost3 := basic.NewPost("judul artikel 3", "lorem ipsum 3", *myuser)
	mypost3.EditContetnt("isi content telah diganti")
	fmt.Println(mypost3.Content)
	fmt.Println(mypost3)
	fmt.Println(mypost3.Owner.GetFullName())

	// membuat struct yang menggunakan interface
	cat := basic.Cat{}
	dog := basic.Dog{}
	basic.Animal(cat)
	basic.Animal(dog)

	// mengakses deffer dan exit
	basic.TryDeffer()
	basic.Tryexit()
	fmt.Println("code setelah exit tidak akan dieksekusi")

}
