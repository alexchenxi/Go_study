package main

import "fmt"

func f3_26() {
MY_LABEL:
	for i := 0; i < 10; i++ {
		if i >= 5 {
			break MY_LABEL
		}
		fmt.Printf("i: %v\n", i)
	}
	fmt.Println("End...")

}

func f2_26() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		// break
		fallthrough // 继续执行第一个case，不做判断
	case 3:
		fmt.Println("3")
		break
	default:
		fmt.Println("default")
	}
}

func f1_26() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			break
		}

	}
}

func test_break() {
	f3_26()
}
