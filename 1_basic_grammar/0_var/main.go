package main

import (
	"fmt"
)

var (
	num int = 5
)

const (
	NUM_ZERO	int = iota		// 0
	NUM_ONE						// 1
	NUM_TWO						// 2
	NUM_THREE					// 3
	NUM_FOUR					// 4
)

func main() {
	fmt.Println("var test...")

	//////////////////////////////////[1. 变量-var]//////////////////////////////////
	//1.1 变量声明
	// var age int				//声明不初始化
	// var age int = 25			//声明并初始化
	// var age = 25				//自动推到类型

	// var age1, age2 = 25, 23		//多变量声明
	// fmt.Printf("[%d] [%d] [%d]\n", age, age1, age2)

	// var (
	// 	name = "James"	//值为James
	// 	ages = 34		//值为34
	// 	salary int		//值为0
	// )
	// fmt.Printf("[%d] [%s] [%d]\n", ages, name, salary)

	//1.2 速记声明
	// no := 23
	// name, salary := "James", 4000
	// fmt.Printf("[no: %d] [name: %s] [salary: %d]\n", no, name, salary)

	fmt.Println(num)
	num := func0()
	fmt.Printf("[%d] [%s]\n", num, "str")
	func2()

	//1.3 匿名变量
	// _, language := func1()
	// fmt.Println(language)

	//////////////////////////////////[2. 常量-const]//////////////////////////////////
	//2.1 常量声明
	// const var_const int = 10
	// const a, b, c = 3, 5, "Gogogo"
	// const (
	// 	size int64 = 1024
	// 	res = 0.0
	// )
	// // var_const = 5	//Error: 不允许对常量进行修改
	// fmt.Printf("[var_const: %d] [a: %d] [b: %d] [c: %s] [size: %d] [res: %f]\n", var_const, a, b, c, size, res)

	//2.2 iota枚举
	// fmt.Printf("[%d] [%d] [%d] [%d] [%d]\n", NUM_ZERO, NUM_ONE, NUM_TWO, NUM_THREE, NUM_FOUR)
}

func func0() int {
	return 1
}

func func1() (int, string) {
	return 1, "Go"
}

func func2() () {
	fmt.Println(num)
}