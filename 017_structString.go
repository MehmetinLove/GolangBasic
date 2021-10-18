package main

import "fmt"

type kisi struct {
	isim string
	soyisim string
	dogumYeri string
	yas int
}

func main()  {
	var bilgilerim kisi
	bilgilerim.isim="MEHMET"
	bilgilerim.soyisim="AŞIK"
	bilgilerim.dogumYeri="KOZAN"
	bilgilerim.yas=45
	fmt.Println(bilgilerim)

}


