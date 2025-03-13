package variable

import (
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

func ShowFunction() {
	showHello()

	result := plusNumber(10, 20)
	println("Result: ", result)

	//Anonymous function
	isNumber := func(n interface{}) bool { // interface{} คือ type ที่รับค่าอะไรก็ได้
		return reflect.TypeOf(n).Kind() == reflect.Int
	}

	result2 := isNumber("Hello")
	println("Result2: ", result2)
}
