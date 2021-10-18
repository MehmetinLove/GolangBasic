package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es")) //test içerisinde es arar bulunursa sonuç true olur.
    fmt.Println(strings.Compare("MEHMET","MeHMET"))// iki string karşılaştırılır. Aynı ise 0, farklı ise -1 döner.
	fmt.Println(strings.Count("ADANA","A"))// ilk string içerisinde kaç tane 2. string olduğunu bulur.
	fmt.Println(strings.HasPrefix("TÜRKİYE CUMHURİYETİ","T"))// baştan harfler kontol edilir doğru ise true, değilse false döner.
}


