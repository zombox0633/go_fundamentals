package variable

import "fmt"

func ShowMap() {
	var priceMap map[string]int64 // map จะเก็บค่าในรูปแบบ key-value

	priceMap = map[string]int64{"A": 100, "B": 200, "C": 300}
	fmt.Printf("Original Map: %v\n", priceMap)

	priceMap["D"] = 400 // เพิ่มข้อมูลใน Map
	fmt.Printf("add Data Map: %v\n", priceMap)

	priceMap["B"] = 150 // แก้ไขข้อมูลใน Map
	fmt.Printf("After Modify Map: %v\n", priceMap)

	delete(priceMap, "C") // ลบข้อมูลใน Map
	fmt.Printf("After Delete Map: %v\n", priceMap)

	value1 := priceMap["A"] // ดึงข้อมูลใน Map
	fmt.Printf("Value of A: %v\n", value1)

	stringMap := map[string]string{"A": "Apple", "B": "Banana", "C": "Cat"}
	fmt.Printf("String Map: %v\n", stringMap)
}
