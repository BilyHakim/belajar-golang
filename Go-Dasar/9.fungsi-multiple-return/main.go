package main

import "math"

// Penerapan Fungsi Multiple Return
//func main() {
//	var diameter float64 = 15
//	var area, circumference = calculate(diameter)
//
//	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
//	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
//}
//
//func calculate(d float64) (float64, float64) {
//	// Hitung Luas
//	var area = math.Pi * math.Pow(d/2, 2)
//	// Hitung Keliling
//	var circumference = math.Pi * d
//
//	// Kembalikan 2 nilai
//	return area, circumference
//}

// Fungsi Dengan Predefined Return Value
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}