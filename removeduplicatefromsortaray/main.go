package main

import "fmt"

func main() {
	sa := []int{1, 1, 2}

	k := removeDuplicates(sa)

	fmt.Println(k)
}

func removeDuplicates(nums []int) int {
	l := 0
	if len(nums) == 0 {
		return l
	}

	for i := 1; i < len(nums); i++ {
		if nums[l] != nums[i] {
			l++
			nums[l] = nums[i]
		}
	}
	fmt.Println(nums)
	return l + 1
}
