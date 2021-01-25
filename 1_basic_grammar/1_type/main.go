package main

import (
	"fmt"
)

func main() {
	//////////////////////////////////[类型-type]//////////////////////////////////
	//1. 类型转换
	// var no int
	// var no_float float32
	// //no_float = no				//Error: 不能给float32类型的变量用int类型赋值
	// no_float = float32(no)		//正确，显示转换
	// fmt.Println(no_float)

	//2. 各类型与其0值
	var var_int int
	var var_bool bool
	var var_string string
	var var_float float64

	fmt.Printf("[%d] [%v] [%s] [%f]\n", var_int, var_bool, var_string, var_float)
}