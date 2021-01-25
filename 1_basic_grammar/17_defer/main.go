package main

import (
	"fmt"
)

func main() {
	Test_deferOrder()
}

func Test_deferOrder() int {
	defer deferTest()
	return getInt()
}

func getInt() int {
	fmt.Println("GetInt")
	return 5
}

func deferTest() {
	fmt.Println("DeferTest")
}