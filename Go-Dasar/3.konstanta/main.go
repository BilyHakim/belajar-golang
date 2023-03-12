package main

//func main() {
//	const firstName string = "john"
//	const lastName = "wick"
//
//	fmt.Print("halo ", firstName, "!\n")
//	fmt.Print("nice to meet you ", lastName, "!\n")
//
//	fmt.Println("john wick")
//	fmt.Println("john", "wick")
//
//	fmt.Print("john wick\n")
//	fmt.Print("john ", "wick\n")
//	fmt.Print("john", " ", "wick\n")
//}

// Deklarasi Multi Konstanta
func main() {
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	// multiple konstanta
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}