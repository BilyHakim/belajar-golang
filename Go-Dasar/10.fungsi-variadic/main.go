package main

import (
	"fmt"
	"strings"
)

//func main() {
//	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
//	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
//	fmt.Println(msg)
//}
//
//func calculate(numbers ...int) float64 {
//	var total int = 0
//	for _, number := range numbers {
//		total += number
//	}
//
//	var avg = float64(total) / float64(len(numbers))
//	return avg
//}

// Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
//func main() {
//	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
//	var avg = calculate(numbers...)
//	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
//
//	fmt.Println(msg)
//}
//
//func calculate(numbers ...int) float64 {
//	var total int = 0
//	for _, number := range numbers {
//		total += number
//	}
//
//	var avg = float64(total) / float64(len(numbers))
//	return avg
//}
//
//var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
//var avg = calculate(numbers...)
//
//// atau
//
//var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)

// Fungsi Dengan Parameter Biasa & Variadic
func main() {
	yourHobbies("wick", "sleeping", "eating")
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s \n", hobbiesAsString)
}