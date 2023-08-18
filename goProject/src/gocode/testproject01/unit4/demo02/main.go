package main

import "fmt"

func main() {
	var count int = 20
	// 单分支：
	if count < 30 {
		fmt.Println("余额不足")
	} else {
		fmt.Println("余额充足")
	}
}
