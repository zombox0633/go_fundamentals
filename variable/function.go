package variable

func showHello() {
	println("Hello World")
}

// go ไม่สามารถกำหนด default value ให้กับ parameter ได้
func plusNumber(a int, b int) int {
	value := a + b
	return value
}

func ShowFunction() {
	showHello()

	result := plusNumber(10, 20)
	println("Result: ", result)
}
