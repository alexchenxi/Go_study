package main

import "fmt"

func main() {
	//变量的声明
	var age int
	//变量的赋值
	age = 18
	//变量的使用
	fmt.Println("age =", age)

	//变量的声明和赋值可以合成一句
	var age2 int = 19
	fmt.Println("age2 =", age2)

	var num int = 12.56
	fmt.Println(num)
	/*
	   # command-line-arguments
	   	.\main.go:17:16: cannot use 12.56 (untyped float constant) as int value in variable declaration (truncated)
	*/
}
