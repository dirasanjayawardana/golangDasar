package basic

import "fmt"

func VariabelDanTipeData() {

	// manifest type (langsung menyebutkan tipe data ketika membuat variabel)
	// type interface (tipe variabel tergantung dari nilainya)

	// variabel yang sudah dibuat harus digunakan, jika tidak maka akan ada warning error
	// format membuat variabel --> var/const namaVariabel tipeData = value
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
	// fmt.Printf("hasil print : % \ns", lastName) // menggunakan print format (%s -> string, %d -> digit, %v --> untuk semua)
	// const string1 = fmt.Sprintf("%s %s", firstName, lastName) // membuat string baru dengan string template

	// const --> variabel yang tidak bisa diubah lagi isinya, wajib langsung diisi ketika variabel dibuat, tetapi tidak ada warning error meski tidak digunakan
	// const contohConstant = "dira sanjaya"
	// fmt.Println("nama saya adalah : " + contohConstant)
	// fmt.Printf("nama saya adalah : %s \n", contohConstant) // menggunakan print format (%s -> string, %d -> digit, %v --> untuk semua)

	// array --> berisi sekumpulan data yg  bertipe data sama, jumlah panjang array ditentukan saat pembuatan, array bersifat pass by value
	// var contohArray = [panjangArray]string{"value1", "value2", "value3"}
	// var contohArray [panjangArray]string
	// contohArray[0] = "value1"
	// contohArray[1] = "value2"

	// slice --> berisi sekumpulan data yg bertipe data sama, panjangnya bisa berubah-ubah sesuai isinya, slice bersifat pass by reference
	// var contohSlice = []string{"value1", "value2", "value3"}
	// contohSlice[0:2] --> contohSlice[indexAwal:indexAhir - 1]

	// map --> berbentuk kay-value, mirip seperti dictionary atau object
	// contohMap:= make(map[tipeDataKey]tipeDataValue)
	person := make(map[string]string)
	person["name"] = "dira"
	person["age"] = "24"
	person["address"] = "jakarta"

	// membuat map dengan langsung diisikan
	student := map[string]string{
		"nama":   "sanjaya",
		"alamat": "bandar lampung",
	}
	fmt.Println(person, "\n", student)
}
