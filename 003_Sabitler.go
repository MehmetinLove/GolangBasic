package main
import "fmt"
func main() {
	const katsayi float64=2 //Değeri daha sonra değiştirilemeyen değişken.

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)// data girişi bu şekilde alınabilir.
	output := input * katsayi
	fmt.Println(output)
}


