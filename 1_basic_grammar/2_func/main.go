package main

import (
	"fmt"
)

type CB func(info string)

func Show(info string) {
	fmt.Printf("[Show]: [%s]\n", info)
}

func GetTitle(year int, name string) string {
	return fmt.Sprintf("[%d NBA Final MVP]: [%s]", year, name)
}

func FuncName(param int) int {
	return 1
}

func GetTotal(price, volume int) int {
	return price * volume
}

func GetPrice() (string, int) {
	return "T-Shirt", 555
}

func GetInfo(width, height int) (area, perimeter int) {
	area = width * height				//具名返回值，无需再声明此变量
	perimeter = (width + height) * 2
	return
}

func main() {
	a, b := GetInfo(5, 5)
	fmt.Printf("[Area: %d] [Perimeter: %d]\n", a, b)

	var cb CB
	cb = Show
	cb("This is a Callback")

	cb2 := GetTitle
	title := cb2(2016, "LeBron James")
	fmt.Println(title)


	// defer + lambda
	defer func(info, info2 string) {
		fmt.Printf("[%s]: [%s]\n", info, info2)
	}("Second defer", "This is a lambda")

	func(info string) {
        fmt.Printf("[Info]: [%s]\n", info)
	}("This is a Anonymous Function")
	
	defer fmt.Println("[First defer] [I'm the First defer]")
}