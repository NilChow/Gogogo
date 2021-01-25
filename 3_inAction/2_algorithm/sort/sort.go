package sort

// 冒泡排序
func Sort_bubble(nums []int) {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
func Sort_selection(nums []int) {
	for i := 0; i < len(nums) - 1; i++ {
		minIdx := i
		for j := i; j < len(nums) - 1; j++ {
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
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}