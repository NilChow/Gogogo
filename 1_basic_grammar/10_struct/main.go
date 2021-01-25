package main

import (
	"fmt"
)

func main() {
	// //结构体的声明
	// type Person struct {
	// 	First_name 	string
	// 	Last_name 	string
	// 	Age 		int			// 开头为小写，不能被其它的包使用
	// 	Sex			bool
	// }
	// type Person struct {
	// 	First_name , last_name 	string
	// 	Age 					int
	// 	sex						bool
	// }

	// //匿名结构体
	// person_1 := struct {
	// 	name 	string
	// 	Age		int
	// 	sex		bool
	// }{"Robinho", 25, true,}
	// fmt.Println(person_1)		//输出：{Robinho 25 true}
	// fmt.Println("===========================================")

	// //声明结构体变量1
	// var person_2 Person
	// person_2.Age = 26
	// person_2.First_name = "Robinho"
	// person_2.last_name = "Chow"
	// person_2.sex = true
	// fmt.Println(person_2)		//输出：{Robinho Chow 26 true}
	// fmt.Println("===========================================")

	// //声明结构体变量2
	// //通过指定字段的名称和对应的值来创建一个结构体，顺序可以不按照结构体声明的顺序来
	// person_3 := Person{
	// 	Last_name	:	"Iverson",
	// 	First_name 	: 	"Allen",
	// 	sex			:	true,
	// 	Age			:	25,
	// }
	// fmt.Println(person_3)		//输出：{Allen Iverson 25 true}
	// fmt.Println("===========================================")

	// //声明结构体变量3
	// //也可以省略字段名称，但是值的顺序必须与结构体声明时字段的顺序保持一致
	// person_4 := Person{"Allen","Iverson",25,true}
	// fmt.Println(person_4)		//输出：{Allen Iverson 25 true}
	// fmt.Println("===========================================")

	// //初始化一部分、忽略一部分(必须用 字段+值 的形式)
	// person_5 := Person{
	// 	First_name	:	"Rudy", 
	// 	last_name	:	"Gay",
	// 	sex			:	true,
	// }
	// fmt.Println(person_5)		//输出：{Rudy Gay 0 true}
	// fmt.Println("===========================================")

	// //结构体指针
	// p_person := &Person{"LeBron", "James", 33, true}
	// fmt.Println("Name:", (*p_person).First_name, (*p_person).last_name )
	// fmt.Println("Age:", (*p_person).Age )
	// fmt.Println("===========================================")

	// fmt.Println("Name:", p_person.First_name, p_person.last_name )
	// fmt.Println("Age:", p_person.Age )
	// fmt.Println("===========================================")

	// //匿名字段
	type Club struct {
		string	// 名字
		int		// 冠军数量
	}
	// var rng Club
	// rng.int = 8
	// rng.string = "RNG"
	// fmt.Println(rng)	//输出：{RNG 8}
	// fmt.Println("===========================================")

	// //结构体嵌套
	type Player struct{
		Name	string
		Age		int
		club	Club
	}
	uzi := Player{
		Name	: "Uzi", 
		Age 	: 23, 
		club 	: Club {"RNG", 8},
	}
	uzi.club.int = 9			// 如果要使用被嵌套的结构体中的字段，必须这样写
	fmt.Println(uzi)			// 输出结果为：{Uzi 23 {RNG 8}}

	// //提阶字段
	type Address struct{
		City		string
		Province	string
	}
	type Brand struct {
		Name		string
		Industry	string
		Address
	}
	laoganma := Brand{}
	laoganma.Name = "LaoGanMa"
	laoganma.Industry = "Food"
	laoganma.City = "GuiYang"
	laoganma.Province = "GuiZhou"
	fmt.Println(laoganma)			// 输出结果为：{LaoGanMa Food {GuiYang GuiZhou}}


	// 结构体的初始值
	player := Player{}
	fmt.Println(player)				// 输出结果为：{ 0 { 0}}

	
	// 结构体比较
	player1 := Player{"", 23, Club{"RNG", 8}}
	player2 := Player{"", 23, Club{"RNG", 8}}
	if player1 == player2 {
		fmt.Println("player1 == player2")        // 此句代码被执行
	}
}