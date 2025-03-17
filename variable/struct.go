package variable

import "fmt"

type EmployeeType struct { // struct เป็นการประกาศ type ให้ข้อมูลรูปแบบ Object คล้ายกับใน JavaScript struct เป็นรูปแบบของ Type ที่ใช้เก็บข้อมูลแบบกำหนดโครงสร้างได้ (คล้าย class แต่ไม่มี method แบบ OOP)
	ID       int
	Name     string
	Position string
}

func ShowStruct() {
	employee1 := EmployeeType{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
	}

	employee2 := EmployeeType{
		ID:       2,
		Name:     "Ten",
		Position: "Ten",
	}

	fmt.Println("Employee ID : ", employee1.ID)

	employeeSlice := []EmployeeType{}
	// employeeSlice[0] = employee1 // ใช้ได้กับ array กับ map เทานั้น
	employeeSlice = append(employeeSlice, employee2) //append ใช้กับ slice เท่านั้นนะจ๊ะ

	fmt.Println("Employee : ", len(employeeSlice))
}
