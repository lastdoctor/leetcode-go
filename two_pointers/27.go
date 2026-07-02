package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	r := removeElement(nums, 2)
	fmt.Println(r, nums)
}
func removeElement(nums []int, val int) int {
	lastIndex := 0
	// [0,1,2,2,3,0,4,2], val = 2
	for currIndex := 0; currIndex < len(nums); currIndex++ {
		candidate := nums[currIndex]
		if candidate == val {
			continue
		}
		nums[lastIndex] = candidate
		lastIndex++
	}

	return lastIndex
}
