package main

import "fmt"

//type student struct {
//	name  string
//	grade int
//}
//
//func (s student) sayHello() {
//	fmt.Println("halo", s.name)
//}
//
//func (s student) getName(i int) string {
//	return strings.Split(s.name, " ")[i-1]
//}
//
//func main() {
//	var s1 = student{"bily hakim", 22}
//	s1.sayHello()
//
//	var name = s1.getName(2)
//	fmt.Println("nama panggilan :", name)
//}

// Method Pointer
type student struct {
	name  string
	grade int
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"bily hakim", 22}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name)
}
