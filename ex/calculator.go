package ex

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) // bufio.NewReader สร้าง buffer สำหรับรับ input จาก console

func GetInputs(prom string) float64 {
	fmt.Printf("Enter %s: ", prom)
	input, _ := reader.ReadString('\n') // รับค่าจาก console และเก็บในตัวแปร input

	// แปลงค่าที่รับมาเป็น float64 โดยตัว strconv คือ package ที่ใช้ในการแปลงค่า
	// ตัว strings คือ package ที่ใช้ในการจัดการ string โดยใช้ฟังก์ชัน TrimSpace ในการลบช่องว่างที่อยู่หน้าและหลัง string
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		message := fmt.Sprintf("%v must be a number only", prom) // Sprintf ใช้ในการสร้าง string โดยใช้ format
		panic(message)                                           // panic ใช้ในการสร้าง error และหยุดการทำงานของโปรแกรม
	}

	return value
}

func getOperator() string {
	fmt.Printf("Enter operator (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func Calculator() {
	num1 := GetInputs("first number")
	num2 := GetInputs("second number")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic("Invalid operator")
	}

	fmt.Printf("Result: %.2f\n", result)
}
