package main

import "fmt"

func test_if() {

	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	// age的作用域只在if语句李

	/* 	a := 1
	   	b := 2
	   	if a > b {
	   		fmt.Printf("a: %v\n", a)
	   	} else {
	   		fmt.Printf("b: %v\n", b)
	   	} */
}
