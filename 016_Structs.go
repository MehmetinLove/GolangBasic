package main

import "fmt"

type prizma struct { //prizma isminde struct yapı tanımlanıyor.
	x, y, h float64
}

func main()  {
	var kup prizma //prizma tipinde kup değişkeni tanımlanıyor.

	kup.x=10 //struct yapıya göre kup alt elemanlarına değer atanır.
	kup.y=2
	kup.h=3
	fmt.Println(kup.x,kup.y,kup.h)
	fmt.Println(alanHesapla(kup.x,kup.y,kup.h))
}

func alanHesapla(a float64,b float64,c float64)  float64{
	alan:=a*b*c
	return alan
}





