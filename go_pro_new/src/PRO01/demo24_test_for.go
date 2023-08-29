package main

import "fmt"

func test_for() {
	f3_24()
}

func f1_24() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f2_24() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3_24() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

func f4_24() {
	for {
		fmt.Println("执行中。。。")
	}
}
