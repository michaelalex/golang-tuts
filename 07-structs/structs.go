package structs

type Person struct {
	Id        int
	Firstname string
	Surname   string
	Age       int
}

type Group struct {
	Id     int
	Name   string
	People []Person
}

// func (r *rect) area() int {
//     return r.width * r.height
// }

// func (r rect) perim() int {
//     return 2*r.width + 2*r.height
// }
