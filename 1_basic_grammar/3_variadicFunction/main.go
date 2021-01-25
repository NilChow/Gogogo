package main

import (
	"fmt"
)

func IsNumExist(num int, num_list ...int) bool {
	for _, v := range num_list {
		if num == v {
			return true
		}
	}
	return false
}

func main() {
	flag := IsNumExist(1, 2, 3, 4, 5, 6, 7, 8, 9, 1)
	if flag == true {
        fmt.Println("Num Exist")            //此句被执行
    } else {
        fmt.Println("Num Doesn't Exist")
	}
	
	num_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    flag2 := IsNumExist(10, num_slice...)
    if flag2 == true {
        fmt.Println("Num Exist")
    } else {
        fmt.Println("Num Doesn't Exist")    //此句被执行
    }
}