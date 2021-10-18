package main
import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)
func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename) //dosya açılır.
	if err != nil {
		return 0, err
	}
	defer f.Close() // açılan dosya en son kapatılır.
	h := crc32.NewIEEE() // şifreleyici oluşturulur
	_, err = io.Copy(h, f)// şifreleme sistemi dosyaya kopyalanır. Yani dosyaya ait bir şifre oluşturulur.
	if err != nil {  //dosyaya ait şifre oluşturma işlemi hatasız yapılması kontrol edilir.
		return 0, err
	}
	return h.Sum32(), nil
}
func main() {
	h1, err := getHash("test1.txt")// birinci dosya şifresi alınır.
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")//2. dosya şifresi alınır.
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)// şifreler ve eşit olup olmadıkları yazdırılır.
}