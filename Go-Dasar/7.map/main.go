package main

//func main() {
//	var chicken map[string]int
//	chicken = map[string]int{}
//
//	chicken["januari"] = 50
//	chicken["februari"] = 40
//
//	fmt.Println("januari", chicken["januari"])
//	fmt.Println("mei", chicken["mei"])
//}

// Inisialisasi Nilai Map
//func main() {
//	var data map[string]int
//	// Akan muncul Error
//	//data["one"] = 1
//
//	// Tidak Error
//	//data = map[string]int{}
//	//data["one"] = 1
//}

//func main() {
//	// Cara Horizontal
//	var chicken1 = map[string]int{"januri": 50, "februari": 40}
//
//	// Cara Vertikal
//	var chicken2 = map[string]int{
//		"januari":  50,
//		"februari": 40,
//	}
//
//	var chicken3 = map[string]int{}
//	var chicken4 = make(map[string]int)
//	var chicken5 = *new(map[string]int)
//}

// Iterasi Item Map Menggunakan for - range
//func main() {
//	var chicken = map[string]int{
//		"januari":  50,
//		"februari": 40,
//		"maret":    34,
//		"april":    67,
//	}
//
//	for key, val := range chicken {
//		fmt.Println(key, " \t:", val)
//	}
//}

// Menghapus Item Map
//func main() {
//	var chicken = map[string]int{"januari": 50, "februari": 40}
//
//	fmt.Println(len(chicken))
//	fmt.Println(chicken)
//
//	delete(chicken, "januari")
//
//	fmt.Println(len(chicken))
//	fmt.Println(chicken)
//}

// Deteksi Keberadaan Item Dengan Key Tertentu
//func main() {
//	var chicken = map[string]int{"januari": 50, "februari": 40}
//	var value, isExist = chicken["mei"]
//
//	if isExist {
//		fmt.Println(value)
//	} else {
//		fmt.Println("item is not exists")
//	}
//}

// Kombinasi Slice & Map
//func main() {
//	var chicken = []map[string]string{
//		map[string]string{"name": "chicken blue", "gender": "male"},
//		map[string]string{"name": "chicken red", "gender": "male"},
//		map[string]string{"name": "chicken yellow", "gender": "female"},
//	}
//	for _, chicken := range chicken {
//		fmt.Println(chicken["gender"], chicken["name"])
//	}
//
//	// Penulisan Go terbaru
//	var chickens = []map[string]string{
//		{"name": "chicken blue", "gender": "male"},
//		{"name": "chicken red", "gender": "male"},
//		{"name": "chicken yellow", "gender": "female"},
//	}
//
//	// Jika key nya berbeda
//	var data = []map[string]string{
//		{"name": "chicken blue", "gender": "male", "color": "brown"},
//		{"address": "mangga street", "id": "k001"},
//		{"community": "chicken lovers"},
//	}
//}
