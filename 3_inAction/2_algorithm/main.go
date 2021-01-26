package main

import (
	"Gogogo/3_inAction/2_algorithm/sort"
	"fmt"
)

func main() {
	fmt.Println("Algorithm Test Start...")

	nums := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	//sort.Sort_bubble(nums)
	//sort.Sort_selection(nums)
	//sort.Sort_insertion(nums)
	sort.Sort_shell(nums)
	//fmt.Println(nums)
}
