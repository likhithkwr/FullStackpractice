package main

import "fmt"

// twoSum returns two different indices whose values add up to target.
// Expected O(n) time and O(n) extra space.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, value := range nums {
		complement := target - value

		// The boolean distinguishes a missing key from an index of zero.
		if earlierIndex, found := seen[complement]; found {
			return []int{earlierIndex, i}
		}

		// Store after checking so the current element cannot match itself.
		seen[value] = i
	}

	// The problem guarantees a pair, so valid inputs return inside the loop.
	return nil
}

// twoSumBruteForce is a simple reference that checks every distinct pair.
// O(n^2) worst-case time and O(1) extra space.
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	examples := []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
		{nums: []int{3, 3}, target: 6},
	}

	for _, example := range examples {
		fmt.Printf("nums=%v, target=%d -> %v\n", example.nums, example.target, twoSum(example.nums, example.target))
	}
}
