package mypackage

import (
	"fmt"
)

var (
	no2 int			// 不能与 mypackage1.go 中的变量重名
)

func init() {
	no2 = 2
	fmt.Println("mypackage2 init function")
}

func Show2() {
	fmt.Printf("This is a function from [mypackage%d]\n", no2)
}