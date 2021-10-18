package main
import "fmt"
func main() {
	defer func() { //normalde bu fonksiyon ertelenir. panic komutu çalştığı için recover ile döndürülür.
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}







