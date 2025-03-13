## Go Module และคำสั่งที่เกี่ยวข้อง

### ลง go
https://go.dev/dl/

### run main
```sh
go run main.go
```

### การสร้าง Go Module
```sh
// go mod init github.com/zombox0633/starter_go
```
เมื่อใช้ `go mod` จะสร้างไฟล์ `go.mod` ขึ้นมา ซึ่งจะทำหน้าที่เหมือน `package.json` ใน Node.js

### การจัดการ Dependencies
```sh
// go mod tidy
```
คำสั่งนี้จะลบแพ็กเกจที่ไม่ได้ใช้ออกจากโปรเจกต์

### การรันและคอมไพล์โปรเจกต์
```sh
// การ run project ใช้คำสั่ง
 go run main.go

// การ build project ใช้คำสั่ง
 go build
```
หลังจาก `go build` จะได้ไฟล์ `main.exe` (Windows) หรือ `main` (Linux/macOS) ซึ่งเป็นไฟล์ไบนารีที่สามารถรันได้โดยตรง

### การติดตั้งและจัดการแพ็กเกจ
```sh
// go get คือคำสั่งที่ใช้ในการดึง package เข้ามาใช้ใน project
 go get github.com/liudng/dogo
```
คำสั่ง `go get` ใช้สำหรับดาวน์โหลดและเพิ่มแพ็กเกจเข้าไปในโปรเจกต์ (`go.mod`)

### การใช้ `dogo`
- `dogo` เป็น Web Server สำหรับรันโปรเจกต์ Go
- ถ้าไม่ได้ใช้ `go install` แต่ใช้ `go get` แพ็กเกจจะถูกเพิ่มลง `go.mod` แต่ยังไม่มี executable พร้อมใช้งาน
- เมื่อลงแพ็กเกจ `dogo` จะมีไฟล์ `go.sum` เกิดขึ้นเพื่อบันทึกเวอร์ชันของแพ็กเกจที่ใช้อยู่

### การตั้งค่า `dogo`
เมื่อใช้งาน `dogo` จำเป็นต้องสร้างไฟล์ `dogo.json` และกำหนดสคริปต์ในการรันโปรเจกต์

### การรันโปรเจกต์ด้วย `dogo`
```sh
// การใช้คำสั่ง dogo จะทำการ run project ของเรา
 dogo
```
`dogo` จะทำให้โปรเจกต์ทำงานแบบ Hot Reload โดยอัตโนมัติ

