package main

import (
	"fmt"
	"strings"
)

func demo15() {

	// 字符串函数
	s := "Hello world"
	fmt.Printf("len(s): %v\n", len(s))

	arr := strings.Split(s, " ")
	fmt.Printf("arr[0]: %v\n", arr[0])

	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "Hello"))
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))
	fmt.Printf("strings.HasSuffix(s, \"world\"): %v\n", strings.HasSuffix(s, "world"))
	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))

	// 字符串的切片操作
	/* 	s := "Welcome to NY"
	   	a := 2
	   	b := 5

	   	fmt.Printf("s[a]: %v\n", s[a])     //查看字符串索引为a的字符编码 l=108
	   	fmt.Printf("s[a:b]: %v\n", s[a:b]) //查看字符串索引从a到b的字符串 前闭后开
	   	fmt.Printf("s[a:]: %v\n", s[a:])
	   	fmt.Printf("s[:b+1]: %v\n", s[:b+1]) */

	// 转义字符

	/* 	s := "Hollo\tworld"
	   	s1 := "c:\\files\\alex"
	   	fmt.Printf("s: %v\n", s)
	   	fmt.Printf("s1: %v\n", s1) */

	/* 	s := "hello"
	   	print(s + "\n")
	   	print(s) */

	// 字符串连接
	/* 	var buffer bytes.Buffer
	   	buffer.WriteString("tom")
	   	buffer.WriteString(",")
	   	buffer.WriteString("20")
	   	fmt.Printf("buffer.String(): %v\n", buffer.String()) */

	/* 	name := "tom"
	   	age := "20"
	   	s := strings.Join([]string{name, age}, ", ")
	   	fmt.Printf("s: %v\n", s) */

	/* 	s1 := "tom"
	   	s2 := "20"

	   	msg := fmt.Sprintf("%s,%s", s1, s2)
	   	fmt.Printf("msg: %v\n", msg) */
	/* 	msg := s1 + s2

	   	fmt.Printf("msg: %v\n", msg) */

	/* 	var s string = "Hello world"
	   	s1 := "Hello world"
	   	fmt.Printf("s: %v\n", s)
	   	fmt.Printf("s1: %v\n", s1)

	   	s4 := `
	   	line1
	   	line2
	   	line3`

	   	fmt.Printf("s4: %v\n", s4) */
}
