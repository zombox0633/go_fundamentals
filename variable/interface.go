package variable

import (
	"fmt"
	"math"
)

type GeometryInterface interface { // interface สำหรับ go ใช้ในการระบุบ function
	area() float64
	perimeter() float64
}

type RectangleType struct { // type สำหรับ go ใช้ในการระบุบตัวแปร
	width, height float64
}

type CircleType struct {
	radius float64
}

func (r RectangleType) area() float64 {
	return r.width * r.height
}

func (r RectangleType) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c CircleType) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) // math.Pow(c.radius, 2) กำลัง 2
}

func (c CircleType) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g GeometryInterface) {
	switch shape := g.(type) {
	case RectangleType:
		fmt.Println("Rectangle")
		fmt.Printf("Width : %.2f , Height : %.2f", shape.width, shape.height)
	case CircleType:
		fmt.Println("Rectangle")
		fmt.Printf("Radius : %.2f", shape.radius)
	default:
		break
	}
	fmt.Println("Area : ", g.area())
	fmt.Println("Perimeter : ", g.perimeter())

}

func ShowInterface() {
	rectangle := RectangleType{
		width:  5,
		height: 10,
	}

	circle := CircleType{
		radius: 15,
	}

	measure(rectangle)
	measure(circle)
}
