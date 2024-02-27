// go mod init namaFolderProject --> inisiasi project golang
// sama seperti di java, di golang harus ada main function untuk menjalankan code, dalam satu project/module hanya boleh ada satu main function, sehingga jika ingin build harus memastikan dalam satu project hanya boleh ada satu main function
// untuk mode develop bisa menjalankan file golang dengan "go run namaFile.go"
// go build --> untuk build/compile file golang, akan menghasilkan file .exe untuk dijalankan diterminal
// go run namaFile.go --> untuk run file golang secara langsung
// di golang boleh menggunaka ; atau tidak

package main

import "fmt"

func main_NOUSED() {
	fmt.Println(("Hello world!"))
}
