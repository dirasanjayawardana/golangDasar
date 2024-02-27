package basic

import (
	"fmt"
	"os"
)

// defer --> untuk menentukan expresi dijalankan diakhir blok eksekusi setelah ekspresi lain dieksekusi
// exit --> memaksa prgram untuk berhenti

func TryDeffer() {
	defer fmt.Println("sampai jumpa")
	fmt.Println("selamat datang")
}

func Tryexit() {
	defer fmt.Println("sanpai jumpa")
	fmt.Println("selamat datang")
	os.Exit(1) // memaksa program untuk berhenti
}
