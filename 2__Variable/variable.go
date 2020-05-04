package main

//Memanggil fmt dari package golang
import (
	"fmt"
)

//Membuat fungsi main yang akan di jalankan di golang
func main() {
	var firstname string = "Jan Falih"
	tinggal := "Tangerang"

	//Membuat printf
	fmt.Printf("Ini Pake Printf (Saya %s Berasal Dari %s)\n", firstname, tinggal)
	//Perbedaan Dengan Println
	fmt.Println("Ini Pake Println (Saya %s Berasal Dari %s)\n", firstname, tinggal)
	//Perbedaan Dengan Print
	fmt.Print("Ini Pake Print (Saya %s Berasal Dari %s)\n", firstname, tinggal)

	//%s akan diganti dengan data string yang ada di parameter ke 2, 3 dan seterusnya
}
