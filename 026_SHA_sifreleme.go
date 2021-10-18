package main
import (
	"crypto/sha1"
	"fmt"
)
func main() {
	h := sha1.New()
	h.Write([]byte("test")) // test yazısının sha1 şifreleme sistemi ile şifrelenmesi
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

//Sonuç:
//[169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]


