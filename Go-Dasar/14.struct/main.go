package main

// Deklarasi Struct
//type student struct {
//	name  string
//	grade int
//}

// Penerapan Struct
//func main() {
//	var s1 student
//	s1.name = "Bily Hakim"
//	s1.grade = 2
//
//	fmt.Println("name :", s1.name)
//	fmt.Println("grade :", s1.grade)
//}

// Inisialisasi Object Struct
//func main() {
//	var s1 = student{}
//	s1.name = "bily"
//	s1.grade = 2
//
//	var s2 = student{"hakim", 2}
//
//	var s3 = student{name: "erlangga"}
//
//	fmt.Println("student 1 :", s1.name)
//	fmt.Println("student 2 :", s2.name)
//	fmt.Println("student 3 :", s3.name)
//}

// Variabel Objek Pointer
//func main() {
//	var s1 = student{name: "bily", grade: 2}
//
//	var s2 *student = &s1
//	fmt.Println("student 1, name :", s1.name)
//	fmt.Println("student 4, name :", s2.name)
//
//	s2.name = "hakim"
//	fmt.Println("student 1, name :", s1.name)
//	fmt.Println("student 4, name :", s2.name)
//}

// Embedded Struct
//type person struct {
//	name string
//	age  int
//}
//
//type student struct {
//	grade int
//	person
//}
//
//func main() {
//	var s1 = student{}
//	s1.name = "bily"
//	s1.age = 22
//	s1.grade = 2
//
//	fmt.Println("name  :", s1.name)
//	fmt.Println("age   :", s1.age)
//	fmt.Println("age   :", s1.person.age)
//	fmt.Println("grade :", s1.grade)
//}

// Embedded Struct Dengan Nama Property Yang Sama
//type person struct {
//	name string
//	age  int
//}
//
//type student struct {
//	person
//	age   int
//	grade int
//}
//
//func main() {
//	var s1 = student{}
//	s1.name = "bily"
//	s1.age = 22
//	s1.person.age = 23
//
//	fmt.Println(s1.name)
//	fmt.Println(s1.age)
//	fmt.Println(s1.person.age)
//}

// Pengisian Nilai Sub-Struct
//func main() {
//	var p1 = person{name: "bily", age: 22}
//	var s1 = student{person: p1, grade: 2}
//
//	fmt.Println("name  :", s1.name)
//	fmt.Println("age   :", s1.age)
//	fmt.Println("grade :", s1.grade)
//}

// Anonymous Struct
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	var s1 = struct {
//		person
//		grade int
//	}{}
//	s1.person = person{"bily", 22}
//	s1.grade = 2
//
//	fmt.Println("name  :", s1.person.name)
//	fmt.Println("age   :", s1.person.age)
//	fmt.Println("grade :", s1.grade)
//}
//
//// anonymous struct tanpa pengisian property
//var s1 = struct {
//	person
//	grade int
//}{}
//
//// anonymous struct dengan pengisian property
//var s2 = struct {
//	person
//	grade int
//}{
//	person: person{"wick", 21},
//	grade:  2,
//}

// Kombinasi Slice & Struct
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	var allStudents = []person{
//		{name: "Wick", age: 23},
//		{name: "Ethan", age: 23},
//		{name: "Bourne", age: 22},
//	}
//
//	for _, student := range allStudents {
//		fmt.Println(student.name, "age is", student.age)
//	}
//}

// Inisialisasi Slice Anonymous Struct
//func main() {
//	var allStudents = []struct {
//		person string
//		grade  int
//	}{
//		{person: person{"bily", 22}, grade: 2},
//		{person: person{"hakim", 22}, grade: 3},
//		{person: person{"ganteng", 24}, grade: 3},
//	}
//	for _, student := range allStudents {
//		fmt.Println(student)
//	}
//}

// Deklarasi Anonymous Struct Menggunakan Keyword var
// Hanya deklarasi
//var student struct{
//	grade int
//}
//
//// deklarasi sekaligus inisialisasi
//var student = struct {
//	grade int
//}{
//	12,
//}

// Nested Struct
//type student struct {
//	person struct{
//		name string
//		age int
//	}
//	grade int
//	hobbies []string
//}

// Deklarasi Dan Inisialisasi Struct Secara Horizontal
//type person struct {
//	name    string
//	age     int
//	hobbies []string
//}

// Tag property dalam struct
//type person struct {
//	name string `tag1`
//	age int `tag2`
//}

//Type alias
type person struct {
	name string
	age int
}
type People = person

var p1 = Person{"bily", 22}
fmt.Println(p1)

var p2 = People{"hakim", 22}
fmt.Println(p2)