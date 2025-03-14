package condition

import "fmt"

func ShowForLoop() {
	// i := 0
	// for i < 5 {
	// 	println(i)
	// 	i++
	// }

	// for j := 0; j < 5; j++ {
	// 	println(j)
	// }

	var input string
	for {
		// fmt.Scanf("%s", &input) // Scanf รับค่าจาก console หลังจากพิมพ์ค่าและกด Enter ใน console ตัว Scanf จะรับค่า "" + "\n" จึงต้องใช้ Scanln แทน
		//เมื่ีออยู่ใน loop ตัว Scanf ตัว input จะถูกอ่านค่าแต่ตัว \n จะยังอยู่ใน buffer จึงต้องใช้ Scanln แทนจึงเกิดการทำงานอีก 1 รอบ

		fmt.Print("Enter input: ")
		fmt.Scanln(&input) // Scanln รับค่าจาก console และเก็บในตัวแปร input
		fmt.Println("Received: ", input)

		if input == "quit" {
			fmt.Println("🦄 : ", input)
			break // ออกจาก loop
		}
	}
}
