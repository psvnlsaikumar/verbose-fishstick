package easy

// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/
func MinimumOperations(nums []int) int {
	var start = 0
	var operations = 0

	for start < len(nums) {
		if isArrayUniqOptimized(nums[start:]) {
			return operations
		}
		start += 3
		operations++
	}

	return operations
}

func isArrayUniq(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return false
			}
		}
	}
	return true
}

func isArrayUniqV2(nums []int) bool {
	var myMap = createMap(nums)
	for _, num := range nums {
		if myMap[num] > 1 {
			return false
		}
	}

	return true
}

func createMap(nums []int) map[int]int {
	var myMap = make(map[int]int)
	for _, num := range nums {
		myMap[num]++
	}

	return myMap

}

func isArrayUniqOptimized(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}
