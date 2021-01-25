package main

import (
	"Gogogo/3_inAction/2_algorithm/sort"
	"fmt"
)

func main() {
	fmt.Println("Algorithm Test Start...")

	nums := []int{61, 17, 29, 22, 34, 60, 72, 21, 50, 1, 62}
	//sort.Sort_bubble(nums)
	//sort.Sort_selection(nums)
	sort.Sort_insertion(nums)
	fmt.Println(nums)
}