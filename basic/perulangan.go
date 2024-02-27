// di golang kondisi di perulangan tidak pakai tanda kurung
// di golang hanya ada for, tetapi bisa digunakan untuk for, while, foreach

package basic

import "fmt"

func Perulangan(){

	// perulangan for
	for i := 0; i <10 ; i ++{
		fmt.Printf("looping ke %d \n", i)
	}
	
	// perulangan while
	var nilai int = 0
	for nilai < 10{
		fmt.Printf("looping ke %d \n", nilai)
		nilai ++
	}

	// perulangan foreach
	drinks := []string{"tea", "coffee", "juice"}
	for index, value := range drinks{
		fmt.Printf("tersedia minuman %s, index ke %d \n", value, index)
	}
}