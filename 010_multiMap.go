package main
import "fmt"

func main() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"isim":"Hidrojen",
			"durum":"gaz",
		},
		"Be": map[string]string{
			"isim":"Berilyum",
			"durum":"katı",
		},

	}
	if el, ok := elements["Be"]; ok {
		fmt.Println(el["isim"], el["durum"])
	}
}