package main

import "fmt"

type WebSites struct {
	Name string
}

func demo16() {
	// %v variable
	/* 	site := WebSites{Name: "alex"}
	   	fmt.Printf("site: %v\n", site)
	   	fmt.Printf("site: %+v\n", site)
	   	fmt.Printf("site: %#v\n", site)
	   	fmt.Printf("\"hello\": %T\n", "hello") //类型 */

	/* 	b := true
	   	fmt.Printf("b: %t\n", b) // 布尔 */

	i := 0x4E2D
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)

	fmt.Printf("i: %c\n", i)
	fmt.Printf("i: %x\n", 100)

	x := 100
	p := &x
	fmt.Printf("p: %p\n", p)
}
