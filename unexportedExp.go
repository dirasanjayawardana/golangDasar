package main

// contoh unexported function --> tidak bisa diakses di package yang berbeda --> menggunakan huruf kecil diawal
func isOdd(num int)bool  {
	return num%2 != 0
}