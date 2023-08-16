# 第一章

## Golang引入

### 简介

Go（又称Golang）是 Google 开发的计算机编程语言

### 设计初衷
解决开发软件难度大问题
1. 主流的编程语言发展落后于硬件，不能合理利用多核多CPU的优势
2. 软件系统复杂度越来越高，维护成本高，缺乏一个足够简介高效的编程语言
3. C/C++ 编译速度慢，内存泄漏

### 应用领域
1. Go服务器端/游戏软件工程师
2. Go分布式/云计算软件工程师
3. 区块链


## Golang简史
Ken Thompson

## 开发环境搭建
golang.org go1.21.0.windows-amd64.msi

## 常见DOS命令
1. 切换盘符： c: d:
2. 显示详细： dir
3. 改变当前目录：cd xxx
4. .当前目录 ..上级目录
5. 清屏：cls
6. 补全命令：tab按键
7. 创建目录：md xxx/ rd xxx
8. 复制/删除文件 copy/del xx.xx

## 测试SDK环境

## Hello World
1. go build test.go
2. test.exe
or
1. go run test.go


# 第二章

## 变量

### 什么是变量
变量是程序世界里的基本单位

### 变量的使用步骤
1. 声明
2. 复制
3. 使用

## 变量的四种使用形式
	// 第一种变量的使用方式
	var num int = 18
	fmt.Println(num)

	// 第二种：指定变量的类型，但不赋值，使用默认值
	var num2 int
	fmt.Println(num2) //打印默认值

	// 第三种：如何没有写变量的类型，那么根据后面的值进行判定变量的类型（自动类型推断）
	var num3 = "10"
	fmt.Println(num3)

	// 第四种：省略var
	gender := "male"
	fmt.Println(gender)

## 声明多个变量

## 数据类型
基本数据类型 复杂数据类型

## 进制和进制转换

## 整数类型
int8 有符号 1字节 -128-127


