package main
import "fmt"
func main() {
	elements := make(map[string]string)		// oluşturulan dizi içerisnden indekse göre değer bulur.
	elements["H"] = "Hidrojen"
	elements["He"] = "Helyum"
	elements["Li"] = "Lityum"
	elements["Be"] = "Berilyum"
	elements["B"] = "Bor"
	elements["C"] = "Karbon"
	elements["N"] = "Azot"
	elements["O"] = "Aksijen"
	elements["F"] = "Flor"
	elements["Ne"] = "Neon"
	elements["Cl"] = "Klor"
	fmt.Println(elements["Cl"])
}