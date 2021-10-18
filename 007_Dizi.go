package main

import "fmt"

func main() {
	dizi1 := [5]float64{ 98, 93, 77, 82, 83 }
	var dizi2=[5]float64{ 98, 93, 77, 82, 83 }
	var dizi [5]float64
	dizi[0] = 98
	dizi[1] = 93
	dizi[2] = 77
	dizi[3] = 82
	dizi[4] = 83
	var toplam float64
	for i := 0; i < 5; i++ {
		toplam += dizi[i]
	}
	fmt.Println(toplam / 5)
	fmt.Println(dizi1)
	fmt.Println(dizi2)
}