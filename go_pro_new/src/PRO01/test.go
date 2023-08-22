package main

import "fmt"

// email:="test@163.com"
var email = "test@163.com"

func getNameAndAge() (string, int) {
	return "tom", 20
}
func test() {
	// var (
	// 	name string
	// 	age  int
	// 	b    bool
	// )
	// name = "alex"
	// age = 30
	// b = true

	// 短变量声明 := 只能放在函数内部
	// name := "tom"
	// age := 20
	// b := true

	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// fmt.Printf("b: %v\n", b)

	// 匿名变量
	fmt.Println("get name and age :")
	name, age := getNameAndAge()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
