package main

import "fmt"

func main() {
	sa := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	k := removeDuplicates(sa)

	fmt.Println(k)
}

func removeDuplicates(nums []int) int {
	l := 0
	prev := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == prev && count >= 2 {
			continue
		}

		if nums[i] != prev {
			count = 0
		}

		nums[l] = nums[i]
		l++
		count++
		prev = nums[i]
	}
	fmt.Println(nums)
	return l
}
