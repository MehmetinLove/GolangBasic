package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	bs, err := ioutil.ReadFile("test.txt") //Dosya okumanın daha kısa yolu.
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}



