package main

import "fmt"

func main()  {
	for i := 0; i <10 ; i++ {
		if i%2==0 { // sayının 2 ile bölümünden kalanı yani tek ya da çift olduğunu bulur.
			fmt.Println(i," sayı çift")
		} else {
			fmt.Println(i," sayı tek")
		}

	}
}
