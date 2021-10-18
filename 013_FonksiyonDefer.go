package main
import "fmt"

func main() {
	// defer komutu ile birlikte çalışan fonksiyon ertelenir.
	// Kendisinden sonra gelen fonksiyon daha önce çalışır.
	defer birinci()
	ikinci()
}

func birinci() {
	fmt.Println("1.fonksiyon")
}
func ikinci() {
	fmt.Println("2.fonksiyon")
}






