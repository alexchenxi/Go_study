package main

import "fmt"

func f1_22() {
	a, b, c := 1, 2, 3
	if a > b {
		if a > c {
			fmt.Printf("Max number is a: %v", a)
		}
	} else {
		if b > c {
			fmt.Printf("Max number is b: %v", b)
		} else {
			fmt.Printf("Max number is c: %v", c)
		}
	}
}

func f2_22() {
	gender := "男"
	age := 12
	if gender == "男" {
		if age >= 18 {
			fmt.Println("成年男性")
		} else {
			fmt.Println("未成年男性")
		}
	} else {
		if age >= 18 {
			fmt.Println("成年女性")
		} else {
			fmt.Println("未成年女性")
		}
	}
}

func test_if_nest() {
	f2_22()
}
