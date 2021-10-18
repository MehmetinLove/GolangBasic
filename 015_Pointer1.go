package main

import "fmt"

func main() {
x := 5
fmt.Println(x) // x değeri 5
xptr:=&x // x'in adres bilgisi xptr değişkenine kaydedilir.
*xptr=20 // xptr'nin tuttuğu adrese 20 kaydedilir.
fmt.Println(x) // x değeri 20
}





