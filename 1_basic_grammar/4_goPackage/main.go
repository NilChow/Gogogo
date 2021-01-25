package main

import (
	"fmt"									// 如果导入了包而没有使用，那么编译会报错
	timer "time"							// 可以给包另起一个名字，之后使用包里的函数和变量都用这个别名
	"Gogogo/1_basic_grammar/4_goPackage/mypackage"			// 导入自定义包

	_ "github.com/go-sql-driver/mysql"		// 如果导入了包，只是想初始化这个包，而不使用，可以加上下划线，避免编译的错误
)

func main() {
	fmt.Println("goPackaet Start...")

	timer.Sleep(3*timer.Second)

	mypackage.Show1()

	timer.Sleep(3*timer.Second)
	
	mypackage.Show2()
}