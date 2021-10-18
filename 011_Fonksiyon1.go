package main

import "fmt"

func main()  {
	a:=degerdonduren()
	fmt.Println(a)
	degerdondurmeyen()
	sonuc:=alanhesapla(20,15)
	fmt.Println(sonuc)
}

func degerdonduren() int {
	return 42
}

func degerdondurmeyen()  {
	fmt.Println("deger yok")
}

func alanhesapla(kenar1 int32,kenar2 int32) int32 {
	alan:=kenar1*kenar2
	return alan

}

