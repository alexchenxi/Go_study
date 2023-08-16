package main

import "fmt"

// 全局变量，定义在函数外的变量(全局变量不能用省略var的声明方式)
var n7 = 7
var n8 = 8

// 一次性声明
var (
	n9  = 500
	n10 = "net"
)

func main() {
	// 定义在{}中的变量叫局部变量
	// 第一种变量的使用方式
	var num int = 18
	fmt.Println(num)

	// 第二种：指定变量的类型，但不赋值，使用默认值
	var num2 int
	fmt.Println(num2) //打印默认值

	// 第三种：如何没有写变量的类型，那么根据后面的值进行判定变量的类型（自动类型推断）
	var num3 = "10"
	fmt.Println(num3)

	// 第四种：省略var
	gender := "male"
	fmt.Println(gender)

	fmt.Println("------------------------------------------------")
	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var n4, name, n5 = 10, "Alex", 9.9
	fmt.Println(n4, name, n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6, height)

	fmt.Println(n7, n8, n9, n10)
}
