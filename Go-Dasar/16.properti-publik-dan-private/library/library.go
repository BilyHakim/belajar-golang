package library

import "fmt"

//func SayHello(name string) {
//	fmt.Println("hello")
//	introduce(name)
//}
//
//func introduce(name string) {
//	fmt.Println("nama saya", name)
//}

// Penggunaan Hak Akses Exported dan Unexported pada Struct dan Propertinya
//type Student struct {
//	Name  string
//	Grade int
//}

// Fungsi init()
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "bily hakim"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
