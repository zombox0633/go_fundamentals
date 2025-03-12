package variable

import "fmt"

func ShowSlice() {
	//การประกาศ Slice โดยไม่กำหนดขนาด
	var priceSlice []int64
	priceSlice = []int64{100, 200, 300} // การกำหนดข้อมูลใน Slice
	fmt.Println("Original Slice:", priceSlice)

	priceSlice = append(priceSlice, 400, 500, 600) // การเพิ่มข้อมูลใน Slice
	fmt.Println("After Append:", priceSlice)

	minPrice := priceSlice[0:3] // การเข้าถึงข้อมูลใน Slice โดยใช้ index 0:3 คือเริ่มต้นที่ 0 และจบที่ 3
	fmt.Println("minPrice Slice:", minPrice)

	minPrice[0] = 50 // การแก้ไขข้อมูลใน Slice จะส่งผลกระทบกับ Slice ต้นฉบับด้วย
	fmt.Println("After modifying minPrice:")
	fmt.Println("minPrice:", minPrice)
	fmt.Println("priceSlice:", priceSlice)

	//ดึงข้อมูล
	fmt.Println("First element:", minPrice[0])
	fmt.Println("Last element:", minPrice[len(minPrice)-1])

	//การลบข้อมูลใน Slice
	minPrice = minPrice[1:]
	fmt.Println("After Delete:", minPrice)
}
