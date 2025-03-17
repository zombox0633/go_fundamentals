package variable

import (
	"fmt"
	"reflect"
)

func showHello() {
	println("Hello World")
}

// go ไม่สามารถกำหนด default value ให้กับ parameter ได้
func plusNumber(a, b int) int {
	value := a + b
	return value
}

func isLoopFor() {
	for i := range 5 {
		defer fmt.Println("i with defer =", i) // defer มีได้มากกว่า 1 ตัวแต่มันจะทำงานเป็น Last in, first out
	}
}

func ShowFunction() {
	// defer จะทำให้การทำงานตามลับปกติที่จากบนลงล่าง กลายเป็นตัวที่มี defer ทำงานเป็นลำดับสุดท้าย
	defer showHello()

	isLoopFor()

	result := plusNumber(10, 20)
	println("Result: ", result)

	//Anonymous function
	isNumber := func(n interface{}) bool { // interface{} คือ type ที่รับค่าอะไรก็ได้
		return reflect.TypeOf(n).Kind() == reflect.Int
	}

	result2 := isNumber("Hello")
	println("Result2: ", result2)
}
