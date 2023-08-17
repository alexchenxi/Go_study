package main

import "fmt"

func main() {
	var age = "hello"
	fmt.Println("age对应的存储空间的地址为：", &age)

	var ptr *string = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为：", *ptr)
}
