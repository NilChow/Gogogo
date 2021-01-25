package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	go Thread_1(nums)
	go Thread_2(nums)
	go Thread_3(nums)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("[Catch Expection: %s]\n", err)
		}
	}()
	time.Sleep(time.Second)
}

func Thread_1(nums []int) {
	for i := range nums {
		fmt.Printf("[Thread_1]: [%d]\n", nums[i])
	}
}

func Thread_2(nums []int) {
	for i := range nums {
		fmt.Printf("[Thread_2]: [%d]\n", nums[i] * 2)
	}
}

func Thread_3(nums []int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("[Catch Expection: %s]\n",err)
		}
	}()

	for i := range nums {
		fmt.Printf("[Thread_3]: [%d]\n", nums[i] * 3)
	}
	fmt.Printf("[Thread_1]: [%d]\n", nums[len(nums)] * 3)
}