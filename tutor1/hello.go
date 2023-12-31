package main

import (
	"fmt"
)

// x เป็นตัวแปร global ที่มีค่าเริ่มต้นเป็น 5
var x int = 5

// y เป็นตัวแปร global โดยไม่ระบุค่าเริ่มต้นอย่างชัดเจน
var y int

func main() {
	// names เป็นตัวแปร string ที่มีค่าที่กำหนดมาล่วงหน้า
	names := "Natthachon Rueangnapharat"

	// แสดงค่าของ x, y, และ names
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(names)

	// เรียกใช้ฟังก์ชัน Init
	Init()

	// ประกาศตัวแปร boolean t และ f
	var t bool = true
	var f bool

	// แสดงค่าของ t และ f
	fmt.Println("t is", t)
	fmt.Println("f is", f)

	// ประกาศตัวแปร age และ favNum พร้อมกำหนดชนิดของตัวแปร
	var age int = 40
	var favNum float64 = 1.6180339

	// แสดงค่าของ age และ favNum
	fmt.Println("age in", age, "favNum in", favNum)

	// ประกาศตัวแปร complex128 เป็นเลข复数
	var complex128 complex128 = 5 + 5i
	fmt.Println("complex128 in", complex128)

	// ประกาศตัวแปร rune r และกำหนดค่าเป็น 10
	var r rune = 10
	fmt.Println("rune is", r)

	// ประกาศตัวแปร string myName
	var myName string = "ณัฐชนม์ เรืองนภารัตน์ "
	fmt.Println(myName + "is a robot")

	// แสดงค่าตัวอักษรที่สามของ myName และความยาวของ myName
	fmt.Println(myName[3])
	fmt.Println("length of myName is", len(myName))

	// ประกาศอาร์เรย์ arry5 ที่มีความยาว 5 และกำหนดค่า
	var arry5 [5]float64
	arry5[0] = 98.7
	arry5[1] = 93.2
	arry5[2] = 77.2
	arry5[3] = 83.7
	arry5[4] = 88.2

	// แสดงค่าของ arry5, ความยาว, และค่าที่ตำแหน่ง 3
	fmt.Println(arry5)
	fmt.Println("length of arry5 is", len(arry5))
	fmt.Println("arry5[3] is", arry5[3])

	// ประกาศอาร์เรย์ arry3 ที่มีความยาว 3 และกำหนดค่า
	arry3 := [3]float64{98, 93, 77}
	fmt.Println(arry3)

	// สร้าง slice s จากสองตัวแรกของ arry5
	var s []float64 = arry5[0:2]
	fmt.Println(s)

	// ประกาศ struct ชื่อ Person
	type Person struct {
		name string
		age  int
	}

	// สร้าง instance ของ Person และกำหนดค่า
	p := Person{name: "ping", age: 21}
	fmt.Println(p)

	// ประกาศตัวแปร x และ xPtr และกำหนดค่า pointer ของ x
	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr)
}

// Init เป็นฟังก์ชันที่แสดงค่าของ x
func Init() {
	fmt.Println("x =", x)
}
