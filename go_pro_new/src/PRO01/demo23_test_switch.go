package main

import (
	"fmt"
)

func test_switch() {
	f4_23()
}

func f1_23() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("一般")
	default:
		fmt.Println("其他")
	}
}

func f2_23() {
	fmt.Println("请输入今天是周几：")
	var day int
	fmt.Scan(&day)
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
	}
}

func f3_23() {
	score := 60
	switch {
	case score >= 60:
		fmt.Println("及格")
	case score >= 80 && score < 90:
		fmt.Println("良好")
	case score >= 90:
		fmt.Println("优秀")
	default:
		fmt.Println("不及格")
	}
}

func f4_23() {
	data_value := 100
	switch data_value {
	case 100:
		fmt.Println("100")
		fallthrough
		// 执行下一个case
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	}
}
