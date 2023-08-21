package main

import "fmt"

func main() {
	const PI = 3.14
	const PI2 = 3.145
	fmt.Printf("PI2: %v\n", PI2)

	const (
		a = 100
		b = 200
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	const w, h = 200, 300
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)

	// iota, 每调用一次增加1
	// const (
	// 	a1 = iota
	// 	a2 = iota
	// 	a3 = iota
	// )

	const (
		a1 = iota
		_
		a2 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	// fmt.Printf("a3: %v\n", a3)

}
