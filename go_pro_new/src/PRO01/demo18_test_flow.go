package main

import (
	"fmt"
)

func flow() {
	a := 100
	if a > 10 {
		fmt.Printf("a: %v\n", a)
	} else {
		fmt.Println("Not enough")
	}

	switch a {
	case 100:
		fmt.Println("a is equal to 100")
	default:
		fmt.Println("default")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	fmt.Println("-----------------------------")

	x := [...]int{1, 2, 3, 7, 8, 9}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
}
