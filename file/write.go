package file

import (
	"os"

	"github.com/zombox0633/starter_go/utils"
)

func writeFile(nameFile string, data []byte) {
	//os.WriteFile จะทำการเขียนไฟล์ทับ แต่ในกรณีที่ไม่มีไฟล์นั้นจะทำการสร้างใหม่ *ไม่ใช่การเพิ่มข้อมูล
	//0644 คือ permission ของไฟล์ | 0644 = ให้เจ้าของไฟล์อ่าน/เขียนได้, คนอื่นอ่านได้แต่แก้ไขไม่ได้
	err := os.WriteFile(nameFile, data, 0644)

	//ทำไมต้องทำเป็น err เพราะเป็น error handling ของ go เมื่อเขียนไม่สำเร็จจะส่ง error
	utils.CheckError(err)
}

func ShowWriteFile() {
	writePet := []byte("Tigger, 2.5, idiot\nToufu, 1.5, cute\nchinchan, 0.7, smart")
	writeFile("./file/data.csv", writePet)

	f, ferr := os.Create("./file/dota2.txt") //os.Create คือคำสั่งสำหรับสร้างไฟล์เปล่า
	utils.CheckError(ferr)
	defer f.Close()

	dota2 := []byte("Abaddon\nBane")
	writeFile("./file/dota2.txt", dota2)
}
