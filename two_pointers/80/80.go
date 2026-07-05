package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	l := removeDuplicates(nums)
	fmt.Println(nums[:l])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	prevIndex := 1
	for currIndex := 2; currIndex < len(nums); currIndex++ {
		candidate := nums[currIndex]
		prevPrevValue := nums[prevIndex-1]
		if prevPrevValue == candidate {
			continue
		}
		prevIndex++
		nums[prevIndex] = candidate
	}

	return prevIndex + 1
}
