package mypackage

import (
	"fmt"
)

var (
	no1 int			// 小写开头，不能被别的包直接使用
)

func init() {
	no1 = 1
	fmt.Println("mypackage1 init function")
}

func Show1() {		// 大写开头，可以被别的包直接使用
	fmt.Printf("This is a function from [mypackage%d]\n", no1)
}