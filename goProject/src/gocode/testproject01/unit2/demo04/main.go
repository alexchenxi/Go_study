package main

import "fmt"

func main() {
	//定义浮点类型的数据：
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = 314E-2
	fmt.Println(num2)

	//浮点数可能会有精度的损失，所以通常情况下，建议使用float64
	var num7 float32 = 256.000000916
	fmt.Println(num7)
	var num8 float64 = 256.000000916
	fmt.Println(num8)

	//golang中默认的浮点类型为float64
	num9 := 3.14
	fmt.Printf("num9对应的默认类型为：%T", num9)
}
