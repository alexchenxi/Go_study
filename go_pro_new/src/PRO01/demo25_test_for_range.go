package main

import "fmt"

func main() {
	f4_25()
}
func f1_25() {
	var a = [...]int{12, 24, 36}

	/* 	for i, v := range a {
		fmt.Printf("i: %v:%v\n", i, v)
	} */

	for _, v := range a {
		fmt.Printf("v:%v\n", v)
	}
}

func f2_25() {
	var s = []int{1, 2, 3}
	for i, v := range s {
		fmt.Printf("%v:%v\n", i, v)
	}
}

func f3_25() {

	var m = make(map[string]any, 0)
	m["name"] = "tom"
	m["age"] = 20
	m["email"] = "test@gmail.com"

	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
}

func f4_25() {
	s := "hello world"
	for _, s := range s {
		fmt.Printf("s: %c\n", s)
	}
}
