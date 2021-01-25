package main

import (
	"fmt"
)

func main() {
	//访问单个字节
	str := "Hello World"
	for i:=0; i<len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// str := "Hello World"
	// for i,v := range str {
	// 	fmt.Printf("[%d] : %c\n", i, v)
	// }

	// // 通过byte切片创建字符串
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xc3, 0xA9}	//这里用对应的十进制数同样奏效
	// str := string(byteSlice)		//这里如果传入数组会报错
	// fmt.Println(str)

	// //潜在隐患--打印输出
	// str := "Señor"
	// for i:=0; i<len(str); i++ {
	// 	fmt.Printf("%c ", str[i])
	// }
	// fmt.Printf("\n")	//输出结果为：S e Ã ± o r

	// //rune解决隐患
	// str := "Señor"
	// str_rune := []rune(str)
	// for i:=0; i<len(str_rune); i++ {
	// 	fmt.Printf("%c ", str_rune[i])
	// }
	// fmt.Printf("\n")	//输出结果为：S e ñ o r

	// //字符串长度
	// str := "Señor"
	// fmt.Println("length of utf-8:", utf8.RuneCountInString(str))	//输出结果为：length of utf-8: 5
	// fmt.Println("length of byte:", len(str))						//输出结果为：length of byte: 6

	// //字符串不可变
	// str := "Hello World"
	// // str[0] = 'Y'        // 非法
	// str = "Yello World"     // 正确
	// fmt.Println(str)

	// //通过rune修改字符串
	// str := "Hello World"
	// new_str := []rune(str)
	// new_str[0] = 'Y'
	// str = string(new_str)
	// fmt.Println(str)

	// //截取字符串
	// sourceStr := "XHelloWorldX"
	// str := sourceStr[1:len(sourceStr)-1]	//输出结果为：HelloWorld (开始的下标，结束的后一位的下标)
	// fmt.Println(str)

	// //截取字符串--中文
	// source := "AB我透你CD"
	// str := string([]rune(source)[2:5])
	// fmt.Println(str)						//输出结果为：我透你
}