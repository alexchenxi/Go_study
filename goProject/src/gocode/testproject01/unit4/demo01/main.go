package main

import "fmt"

func main() {
	var count int = 3
	// 单分支：
	if count < 30 {
		fmt.Println("余额不足")
	}
	fmt.Println(count)

	// if后面一定要有空格
	// {}一定不能省略
	// 条件表达式左右的（）是建议省略的
	// 在golang中，if后面可以并列的加入变量的定义

	if count := 20; count < 30 {
		fmt.Println("not enough")
		fmt.Println(count)
	}
}
