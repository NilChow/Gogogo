package main

import (
	"fmt"
	"time"
)

func main() {
	/////////////// [if-else] ///////////////
	if num:=10; num >= 10 {            //先声明num变量并初始化为10；再判断条件
		fmt.Println("num >= 10")       //num的作用域仅限于这个if-else语句中
	} else {
		fmt.Println("num < 10")
	}

	/////////////// [for] ///////////////
	// For_Test()

	/////////////// [switch-case] ///////////////
	Switch_Test()
}

func For_Test() {
	// for
	for i:=1; i<=10; i++ {
		if i == 4 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Printf("[Times = %d]\n", i)
	}

	// endless for
	count := 1
	for {
		fmt.Printf("[Count = %d]\n", count)
		count++
		time.Sleep(1*time.Second)
	}
}

func Switch_Test() {
	dayOfWeek := 2

	// Switch 1
    switch dayOfWeek {
    case 1:
        fmt.Println("[Switch 1] [It's Mondy]")
    case 2:
        fmt.Println("[Switch 1] [It's Tuesday]")
    case 3:
        fmt.Println("[Switch 1] [It's Wednesday]")
    case 4:
        fmt.Println("[Switch 1] [It's Thursday]")
    case 5:
        fmt.Println("[Switch 1] [It's Friday]")
    case 6:
        fmt.Println("[Switch 1] [It's Saturday]")
    case 0:
        fmt.Println("[Switch 1] [It's Sunday]")
    default:
        fmt.Println("[Switch 1] [Error]")
	}
	
	// Switch 2
	switch dayOfWeek {
    case 1, 2, 3, 4, 5 :
        fmt.Println("[Switch 2] [It's Working Day]")
    case 6, 0:
        fmt.Println("[Switch 2] [It's Weekend]")
    default:
        fmt.Println("[Switch 2] [Error]")
	}

	// Switch 3
	switch {
    case dayOfWeek >= 1 && dayOfWeek <= 5 :
        fmt.Println("[Switch 3] [It's Working Day]")
    case dayOfWeek == 0 || dayOfWeek == 6:
        fmt.Println("[Switch 3] [It's Weekend]")
    default:
        fmt.Println("[Switch 3] [Error]")
	}
	
	// fallthrough
	num := 2
	switch num {
    case 1:
        fmt.Println("[fallthrough] [It's 1]")
    case 2:
		fmt.Println("[fallthrough] [It's 2]")
		fallthrough
    case 3:
		fmt.Println("[fallthrough] [It's 3]")
		fallthrough
    case 4:
        fmt.Println("[fallthrough] [It's 4]")
    case 5:
        fmt.Println("[fallthrough] [It's 5]")
    default:
        fmt.Println("[fallthrough] [Error]")
	}
}