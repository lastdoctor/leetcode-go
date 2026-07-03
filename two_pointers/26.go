package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	prevIndex := 0
	// [0,0,1,1,1,2,2,3,3,4]
	for currIndex := 1; currIndex < len(nums); currIndex++ {
		candidate := nums[currIndex]
		prevValue := nums[prevIndex]
		if candidate == prevValue {
			continue
		}
		prevIndex++
		nums[prevIndex] = candidate
	}

	return prevIndex + 1
}
