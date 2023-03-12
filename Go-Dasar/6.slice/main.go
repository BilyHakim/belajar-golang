package main

import "fmt"

//func main() {
//	var fruits = []string{"apple", "grape", "banana", "melon"}
//	fmt.Println(fruits[0])
//}

//func main() {
//	var fruits = []string{"apple", "grape", "banana", "melon"}
//	var newFruits = fruits[0:2]
//
//	fmt.Println(newFruits)
//}

//func main() {
//	var fruits = []string{"apple", "grape", "banana", "melon"}
//
//	var aFruits = fruits[0:3]
//	var bFruits = fruits[1:4]
//
//	var aaFruits = fruits[1:2]
//	var baFruits = fruits[0:1]
//
//	fmt.Println(fruits)
//	fmt.Println(aFruits)
//	fmt.Println(bFruits)
//	fmt.Println(aaFruits)
//	fmt.Println(baFruits)
//
//	// Buah "grape" diubah menjadi "pinnaple"
//	baFruits[0] = "pinnaple"
//
//	fmt.Println(fruits)   // [apple pinnaple banana melon]
//	fmt.Println(aFruits)  // [apple pinnaple banana]
//	fmt.Println(bFruits)  // [pinnaple banana melon]
//	fmt.Println(aaFruits) // [pinnaple]
//	fmt.Println(baFruits) // [pinnaple]
//}

//func main() {
//	var fruits = []string{"apple", "grape", "banana", "melon"}
//
//	fmt.Println(len(fruits))
//}

//func main() {
//	var fruits = []string{"apple", "grape", "banana", "melon"}
//	fmt.Println(len(fruits)) // len: 4
//	fmt.Println(cap(fruits)) // cap: 4
//
//	var aFruits = fruits[0:3]
//	fmt.Println(len(aFruits)) // len: 3
//	fmt.Println(cap(aFruits)) // cap: 4
//
//	var bFruits = fruits[1:4]
//	fmt.Println(len(bFruits)) // len: 3
//	fmt.Println(cap(bFruits)) // cap: 3
//}

// Fungsi Append
//func main() {
//	var fruits = []string{"apple", "grape", "banana"}
//	var cfruits = append(fruits, "papaya")
//
//	fmt.Println(fruits)
//	fmt.Println(cfruits)
//}

//func main() {
//	var fruits = []string{"apple", "grape", "banana"}
//	var bFruits = fruits[0:2]
//
//	fmt.Println(cap(bFruits)) // 3
//	fmt.Println(len(bFruits)) // 2
//
//	fmt.Println(fruits)  // ["apple", "grape", "banana"]
//	fmt.Println(bFruits) // ["apple", "grape"]
//
//	var cFruits = append(bFruits, "papaya")
//
//	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
//	fmt.Println(bFruits) // ["apple", "grape"]
//	fmt.Println(cFruits) // ["apple", "grape", "papaya"]
//}

// Fungsi Copy
//func main() {
//	dst := make([]string, 3)
//	src := []string{"watermelon", "pinnaple", "apple", "orange"}
//	n := copy(dst, src)
//
//	fmt.Println(dst)
//	fmt.Println(src)
//	fmt.Println(n)
//}

//func main() {
//	dst := []string{"potato", "potato", "potato"}
//	src := []string{"watermelon", "pinnaple"}
//	n := copy(dst, src)
//
//	fmt.Println(dst) // watermelon pinnaple potato
//	fmt.Println(src) // watermelon pinnaple
//	fmt.Println(n)   // 2
//}

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}