package variable

import (
	"fmt"
	"time"
)

// ตัวอย่างการทำงานแบบขนานของ Concurrency หรือ routine ของ go
// โดยการใช้ go นั้นจะทำให้การทำงานของ function นั้นๆ ทำงานพร้อมกัน โดยไม่ต้องรอให้ function ก่อนหน้าทำงานเสร็จก่อน
func IsLoopString(String string) {
	for i := 0; i < 100; i++ {
		fmt.Println(String, ":", i)
	}
}

func processOne(c chan string, text string) {
	c <- text //การส่งจากจาก text ไปใส่ใน c จะทำให้ค่าของ ch มีการเปลี่ยนแปลงค่า
	c <- "10" // การใส่ค่าใหม่จะไม่ทำการเปลี่ยนแปลงค่าเก่า
}

func ShowChannel() {
	// โดยปกติการทำงานจะเป็นแบบบนลงล่าง แต่การใช้งานคำสั่ง go หรือที่เรียก routine จะทำงานแบบ ขนาน แบบทำไปพร้อมๆกัน

	// การทำงานแบบปกติ
	// IsLoopString("🦄")
	// IsLoopString("🐊")

	// การทำงานแบบ routine *การใช้คำสั่ง go จำเป็นต้องใส่ time.Sleep ยังงั้นมันจะไม่ออก
	// go IsLoopString("🐲")
	// go IsLoopString("🦝")
	// go IsLoopString("🐳")
	// time.Sleep(10 * time.Second) //ระยะเวลามีผลกับการทำงานของมัน

	//----------------------------------------------------------------------------------------------------------
	// make ใช้สำหรับสร้าง chan, slice, map | make(chan string, 1) ส่วนสุดท้ายคือ Buffered Channel คือตัวระบุบจำนวนที่ใส่ลง Channel
	ch := make(chan string)  // ใช้คำสั่ง make ในการสร้าง channel สำหรับ string และเป็นแบบ Unbuffered Channel
	go processOne(ch, "Ten") // ใช้ Goroutine ถ้าไม่ใช้จะเกิด deadlock!
	fmt.Println(<-ch)        // สำหรับนำค่า ch มาใช้แต่จะได้ค่า chan แรกที่เก็บ
	fmt.Println(<-ch)        // chan ที่สองเก็บ
	time.Sleep(5 * time.Second)
}
