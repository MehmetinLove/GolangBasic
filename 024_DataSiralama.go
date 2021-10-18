package main
import ("fmt" ; "sort")
type Person struct {
	Name string
	Age int
}
type ByName []Person
func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
type ByAge []Person
func (ps ByAge) Len() int {
	return len(ps)
}
func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}
func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	kids := []Person{
		{"Erdem",17},
		{"Mina",12},
		{"Ayşegül",44},
		{"Mehmet",45},
	}
	sort.Sort(ByName(kids))// Dataları isme göre sıralar.
	fmt.Println(kids)
	sort.Sort(ByAge(kids))// Dataları yaşa göre sıralar.
	fmt.Println(kids)
}


