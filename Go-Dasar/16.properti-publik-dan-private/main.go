package main

import (
	"16.properti-publik-dan-private/library"
	"fmt"
)

//func main() {
//	library.SayHello("bily")
//}

//func main() {
//	var s1 = library.Student{"bily", 22}
//	fmt.Println("name ", s1.Name)
//	fmt.Println("grade ", s1.Grade)
//}

// 26.5. Import Dengan Prefix Tanda Titik
//func main() {
//	var s1 = Student{"bily", 22}
//	fmt.Println("name", s1.Name)
//	fmt.Println("grade", s1.Grade)
//}

// Pemanfaatan Alias Ketika Import Package
//func main() {
//	f.Println("Hello World!")
//}

//func main() {
//	sayHello("bily")
//}

func main() {
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
