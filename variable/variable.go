package variable

import "fmt"

func ShowVariable() {
	// การประกาศตัวแปร
	var x string = "Hello World"
	fmt.Println("test" + x)

	// การประกาศตัวแปรตัวแปรหลายตัวในบรรทัดเดียวกันได้
	var y, z int = 10, 1
	// fmt.Println(z, y)
	fmt.Println(z + y)

	//การประกาศตัวแปรแบบ short hand
	numberFloat := 3.33
	// fmt.Println(numberFloat)
	fmt.Println(numberFloat + float64(y)) // ต้องแปลงชนิดข้อมูลก่อนใช้งาน
}
