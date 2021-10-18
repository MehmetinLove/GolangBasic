package main
import (
	"fmt"
	"os"
)
func main() {
	file, err := os.Open("test.txt")// test.txt dosyası açılır.
	if err != nil {
		file,err=os.Create("test.txt")// Eğer dosya açılmaz ise dosya oluşturulur.
		fmt.Fprintf(file,"deneme") //Dosya içine yazılır
		//file.WriteString("test") //Dosya içine yazmanın farklı yolu
		return
	}
	defer file.Close() // Dosya işlemleri bittiğinde dosya kapatılır.

	stat, err := file.Stat() // Dosya boyutu alınır.
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())// Dosya büyüklüğü kadar string hücre ayrılır.
	_, err = file.Read(bs) // Dosya okunur ve bs içerisine kaydedilir.
	if err != nil {
		return
	}
	str := string(bs)// Okunan string ifade yazdırılır.
	fmt.Println(str)
}



