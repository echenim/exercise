package main

import "fmt"

func main() {
	sa := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(sa)
	k := removeElement(sa, 2)

	fmt.Println(k)
}

func removeElement(nums []int, val int) int {

	count := 0
	for i, num := range nums {
		nums[i-count] = nums[i]
		if num == val {
			count++
		}
	}
	fmt.Println(nums)
	return len(nums) - count
}
