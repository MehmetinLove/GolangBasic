package main

import "fmt"

func  main()  {
// Değşken tanımlama şekilleri
	var yazi1 = "Mehmet"
	var yazi2 string= "ASIK"
	yazi3 :="GO Programlama" // Değişken tipi belirtilmeden ve değişken daha önce tanımlanmadan da kullanılabilir.
	var sayi1, sayi2 int = 1, 2
	var sayi3 float32=5.5
	sayi4:=-15.6
	var lojik bool
	lojik=!true //! işareti ile işlemin tersi alınabilir.
	fmt.Println(yazi1,yazi2,yazi3, sayi1, sayi2,sayi3,sayi4, lojik)
	fmt.Println("1 + 2 =", 1.0 + 2.0)
	fmt.Println("4/3 =", 4.0 / 3.0)
	fmt.Println(yazi1==yazi2) // 2 string ifade karşılaştırılır. Sonuç true ya da false olarak döner.
}
