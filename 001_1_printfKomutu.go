package main

import (
	"fmt"
	"os"
)
type point struct { // struct yapısında değişken tanımlama
	x, y int
}

func main()  {
	p := point{1, 2}
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "ABCD")
	fmt.Printf("%p\n", &p)

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
