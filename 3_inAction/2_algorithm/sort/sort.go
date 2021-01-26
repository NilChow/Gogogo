package sort

import (
	"fmt"
)

// 冒泡排序
func Sort_bubble(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
func Sort_selection(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i; j < len(nums)-1; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		if i != minIdx {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

// 插入排序
func Sort_insertion(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// 希尔排序
func Sort_shell(nums []int) {
	for stepLength := len(nums) / 2; stepLength >= 1; stepLength = stepLength / 2 {
		fmt.Printf("==============================[StepLength: %d]==============================\n", stepLength)
		for i := 0; i < len(nums); i++ {
			fmt.Printf("%d   ", nums[i])
			if (i+1)%stepLength == 0 {
				fmt.Println()
			}
		}
		fmt.Println()

		for i := 0; i < stepLength; i++ {
			for j := i; j+stepLength < len(nums); {
				if nums[j] > nums[j+stepLength] {
					nums[j+stepLength], nums[j] = nums[j], nums[j+stepLength]
				} else {
					j += stepLength
				}
			}
		}

		for i := 0; i < len(nums); i++ {
			fmt.Printf("%d   ", nums[i])
			if (i+1)%stepLength == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		fmt.Println(nums)
	}
}
