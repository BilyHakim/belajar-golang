package main

import "fmt"

//func main() {
//	var names [4]string
//	names[0] = "trafalgar"
//	names[1] = "d"
//	names[2] = "water"
//	names[3] = "law"
//
//	fmt.Println(names[0], names[1], names[2], names[3])
//}

// Pengisian Elemen Array yang Melebihi Alokasi Awal
//func main() {
//	var names [4]string
//	names[0] = "trafalgar"
//	names[1] = "d"
//	names[2] = "water"
//	names[3] = "law"
//	names[4] = "bilgan"
//
//	fmt.Println(names[0], names[1], names[2], names[3])
//}

// Inisialisasi Nilai Awal Array
//func main() {
//	var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	fmt.Println("Jumlah element \t\t", len(fruits))
//	fmt.Println("Isi semua element \t", fruits)
//}

// Inisialisasi Nilai Array Dengan Gaya Vertikal
//func main() {
//	var fruits [4]string
//	// Cara Vertikal
//	//var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	// Cara Horizontal
//	fruits = [4]string{
//		"apple",
//		"grape",
//		"banana",
//		"melon",
//	}
//
//	fmt.Println("Isi semua element \t", fruits)
//}

// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
//func main() {
//	var numbers = [...]int{2, 3, 2, 4, 3}
//
//	fmt.Println("data array \t:", numbers)
//	fmt.Println("jumlah elemen \t:", len(numbers))
//}

// Array Multidimensi
//func main() {
//	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
//	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
//
//	fmt.Println("numbers1", numbers1)
//	fmt.Println("numbers2", numbers2)
//}

// Perulangan Elemen Array Menggunakan Keyword for
//func main() {
//	var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	for i := 0; i < len(fruits); i++ {
//		fmt.Printf("elemen %d : %s\n", i, fruits[i])
//	}
//}

// Perulangan Elemen Array Menggunakan Keyword for - range
//func main() {
//	var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	for i, fruit := range fruits {
//		fmt.Printf("elemen %d : %s\n", i, fruit)
//	}
//}

// Penggunaan Variabel Underscore _ Dalam for - range
// Error
//func main() {
//	var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	for i, fruit := range fruits {
//		fmt.Printf("nama buah : %s\n", fruit)
//	}
//}

// Solve
//func main() {
//	var fruits = [4]string{"apple", "grape", "banana", "melon"}
//
//	for _, fruit := range fruits {
//		fmt.Printf("nama buah : %s\n", fruit)
//	}
//}

// Alokasi Elemen Array Menggunakan Keyword make
func main() {
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits)
}