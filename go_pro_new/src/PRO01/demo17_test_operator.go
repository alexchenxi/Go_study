package main

import "fmt"

func getA() int {
	return 20
}

func operator() {

	a := 100
	a = 200
	fmt.Printf("a: %v\n", a)
	a = getA()
	fmt.Printf("a: %v\n", a)

	a += 100
	fmt.Printf("a: %v\n", a)

	// 位运算 & |
	/* 	a := 4
	   	fmt.Printf("a: %b\n", a)
	   	b := 8
	   	fmt.Printf("b: %b\n", b)

	   	r := a & b
	   	fmt.Printf("r: %v\n", r)
	   	r = a | b
	   	fmt.Printf("r: %v\n", r)
	   	r = a ^ b
	   	fmt.Printf("r: %v\n", r)
	   	r = a << 2
	   	fmt.Printf("r: %v\n", r)
	   	r = a >> 2
	   	fmt.Printf("r: %v\n", r) */

	/* 	a := 100
	   	b := 20

	   	fmt.Printf("(a + b): %v\n", (a + b))
	   	fmt.Printf("(a - b): %v\n", (a - b))
	   	fmt.Printf("(a * b): %v\n", (a * b))
	   	fmt.Printf("(a / b): %v\n", (a / b))
	   	fmt.Printf("(a mod b): %v\n", (a % b)) */

}
