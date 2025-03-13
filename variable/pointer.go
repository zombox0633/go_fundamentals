package variable

import "fmt"

func normalFunction(ivalue int) {
	ivalue = 2 //โดยปกติการรับ argument จะทำการรับค่าโดยไม่เปลี่ยนแปลงค่าตัวแปรต้นทาง
}

func pointerFunction(ipointer *int) { // การรับ argument แบบ pointer จะใส่ * หน้า type ของ argument
	*ipointer = 3 // การรับ argument แบบ pointer จะเปลี่ยนแปลงค่าตัวแปรต้นทางด้วย
}

func ShowPointer() {
	i := 1
	fmt.Println("initial:", i)

	normalFunction(i)
	fmt.Println("normalFunction:", i)

	pointerFunction(&i) // การส่งค่าตัวแปรต้นทางเข้าไปใน function แบบ pointer จะใส่ & หน้าตัวแปร
	fmt.Println("pointerFunction:", i)
}
