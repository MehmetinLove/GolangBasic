package main

import "fmt"

func main()  {
	var slice1[]int // sınırı belli olmayan slice tanımlandı.
	slice2:=make([]int,5) // make ile slice boyutu belirlenir.
	slice3 := make([]int, 5, 10) //slice için 10 birimden 5 tanesi ayrılır.

	slice1 = []int{1,2,3}
	slice2 = append(slice1, 4, 5)	// slice üstüne ekleme yapılır.
	copy(slice3,slice2)						//slice2 slice3 içerisine kopyalanır.

	fmt.Println(slice1, slice2,slice3)
}
