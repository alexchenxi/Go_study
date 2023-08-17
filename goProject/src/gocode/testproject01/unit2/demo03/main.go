package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型
	//var num1 int8 = 24
	//fmt.Print(num1)
	/*
		# command-line-arguments
		.\main.go:7:18: cannot use 224 (untyped int constant) as int8 value in variable declaration (overflows)
	*/

	//var num2 uint8 = -20
	//fmt.Println(num2)
	/*
		# command-line-arguments
		.\main.go:14:19: cannot use -20 (untyped int constant) as uint8 value in variable declaration (overflows)
	*/

	var num3 uint64 = 1000000000000000
	//Printf函数的作用是：格式化的，把num3的类型填充到%T的位置上
	fmt.Printf(`num3的类型是：%T`, num3)
	fmt.Println()
	//unsafe.Sizeof()占用字节数
	fmt.Println(unsafe.Sizeof(num3))


	//表示年龄
	var age byte = 18
	fmt.Println(unsafe.Sizeof(age))
}
