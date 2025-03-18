package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zombox0633/starter_go/utils"
)

func ShowReadFile() {
	f, err := os.Open("./file/data.csv") //ทำการเลือกและเปิดไฟล์
	// defer f.Close()

	utils.CheckError(err)

	scanner := bufio.NewScanner(f) //ทำหน้าที่ในการอ่านไฟล์

	c := 1
	for scanner.Scan() {
		line := scanner.Text()           //เลือกแต่ละ line เลือกจาก \n
		item := strings.Split(line, ",") //ทำให้ข้อมูลอยู่ในรู้แบบ array
		// fmt.Printf("Line %d : %s\n", c, line)

		fmt.Println(item[2])
		c++
	}

}
