package main

import (
	"fmt"
)

type Human struct {
	age		int
	name	string
	sex		string
}

func (p Human) String() string {
	return fmt.Sprintf("[Name: %s] [Age: %d] [Sex: %s]", p.name, p.age, p.sex)
}

func main () {
	// //指针的声明与访问
	// var num int = 5
	// p_num := &num
	// fmt.Printf("[%p] [%d]\n", p_num, *p_num)				// 输出结果为：[0xc0000120a0] [5]
	// fmt.Println("===========================================")

	// p_string := new(string)
	// *p_string = "It's Taco Tuesday, gi gi gi gi gi gi gi"
	// fmt.Println(*p_string)

	// p := &Human{34, "James", "male"}
	// fmt.Println(*p)											// 输出结果为：[Name: James] [Age: 34] [Sex: male]

	// //打印类型名称
	// n := 5
	// p_int := &n
	// fmt.Printf("[%T] [%T]\n", n, p_int)							// 输出结果为：[int] [*int]

	// //空值
	// a := 25
	// var p_a *int
	// if p_a == nil {
	// 	fmt.Println("p_a is nil : ", p_a)
	// 	p_a = &a
	// 	fmt.Println("p_a address is :", p_a, ", value is :", *p_a)
	// }
	// fmt.Println("===========================================")

	// //Go不支持指针运算
	// str := "Golang"
	// p_str := &str
	// fmt.Println(*p_str)
	// p_str++
}