package main

import "fmt"

func main() {

	var num1 = 10
	fmt.Println(num1)
	var num2 = (10+20)%3 + 3 - 7
	fmt.Println(num2)

	var num3 = 10
	num3 += 20
	fmt.Println(num3)

	//交换值并输出
	var a = 8
	var b = 4
	fmt.Printf("a=%v,b=%v", a, b)

}
