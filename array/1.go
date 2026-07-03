package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// [2,7,11,15], target = 9;
	for index, value := range nums {
		candidate := target - value
		mapIndex, ok := m[candidate]
		if ok {
			return []int{index, mapIndex}
		}

		m[value] = index
	}

	return nil
}
