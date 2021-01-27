package sort

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

// 希尔排序（分组的插入排序）
func Sort_shell(nums []int) {
	for stepLength := len(nums) / 2; stepLength >= 1; stepLength = stepLength / 2 {
		for i := 0; i < stepLength; i++ {
			for idx, j := i+stepLength, i+stepLength; j < len(nums); {
				if idx >= stepLength && nums[idx] < nums[idx-stepLength] {
					nums[idx], nums[idx-stepLength] = nums[idx-stepLength], nums[idx]
					idx -= stepLength
				} else {
					j += stepLength
					idx = j
				}
			}
		}
	}
}

// 归并排序
func Sort_merge(nums []int) {

}

// 快速排序
func Sort_quick(nums []int) {
	left, right := 0, len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; {
		for idx := i; idx < len(nums); idx++ {
			if nums[idx] > nums[i] {
				nums[idx], nums[i] = nums[i], nums[idx]
				j = idx
			}
		}

		for idx :=
	}
}