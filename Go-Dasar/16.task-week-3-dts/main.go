package main

import "fmt"

func main() {
	kalimat := "selamat malam"

	jumlahKata := make(map[string]int)

	for _, kata := range kalimat {
		jumlahKata[string(kata)]++
		fmt.Println(string(kata))
	}
	fmt.Println(jumlahKata)
}
