// di golang kondisi di percabangan tidak pakai tanda kurung

package basic

func Percabangan(){

	if true {
		println("eksekusi blok if")
	} else if false {
		println("eksekusi blok else if")
	} else {
		println("eksekusi blok else")
	}
	
}