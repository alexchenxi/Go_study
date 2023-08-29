package main

import "fmt"

func f2() {
	var num int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("%v是偶数", num)
	} else {
		fmt.Printf("%v是奇数", num)
	}
}

func test_if_else() {
	f2()
	/* var name string
	var age int
	var email string
	fmt.Println("请输入name,age,email,用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email) */
}
