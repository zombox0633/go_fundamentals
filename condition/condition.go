package condition

import "fmt"

// go ไม่มี ternary operator ใช้ได้แต่ if else

func checkNumber(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
}

func ShowCondition() {
	number := 3
	result := checkNumber(number)
	fmt.Printf("Number %d is %s\n", number, result)
}
