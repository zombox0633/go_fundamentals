package variable

import "fmt"

func ShowArray() {
	//การประกาศ Array
	var priceArray [3]int64 // การประกาศ Array โดย [3]int คือจำนวนข้อมูลที่สามารถเก็บได้
	priceArray[0] = 100
	fmt.Println(priceArray)

	//การประกาศตัวแปร array แบบ short hand
	stringArray := [3]string{"Hello", "World", "Go"} // การประกาศแบบนี้จะทำให้ array มีขนาดตายตัว
	fmt.Println(stringArray)

}
