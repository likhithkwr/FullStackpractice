package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

// unfinished marks only the intentionally empty practice method.
var unfinished = errors.New("TODO: implement twoSum()")

// twoSum should return the indices of two different elements that add up to target.
// Exactly one valid pair exists, and either index order is accepted.
func twoSum(nums []int, target int) []int {
	// TODO: Replace the placeholder below with your own solution.
	panic(unfinished)
}

type exampleCase struct {
	nums     []int
	target   int
	expected []int
}

type checkResult struct {
	answer []int
	status string
	detail string
}

// The checker validates the returned indices against the original input.
func isValidAnswer(nums []int, target int, answer []int) bool {
	if len(answer) != 2 {
		return false
	}
	first, second := answer[0], answer[1]
	return first >= 0 && first < len(nums) &&
		second >= 0 && second < len(nums) &&
		first != second && nums[first]+nums[second] == target
}

func checkOne(nums []int, target int) (result checkResult) {
	// Report an unfinished method as TODO and other panics as errors.
	defer func() {
		if recovered := recover(); recovered != nil {
			if recovered == unfinished {
				result.status = "TODO"
			} else {
				result.status = "ERROR"
				result.detail = fmt.Sprint(recovered)
			}
		}
	}()

	result.answer = twoSum(slices.Clone(nums), target)
	if isValidAnswer(nums, target, result.answer) {
		result.status = "PASS"
	} else {
		result.status = "FAIL"
	}
	return result
}

func main() {
	cases := []exampleCase{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{nums: []int{0, 4, 3, 0}, target: 0, expected: []int{0, 3}},
		{nums: []int{-3, 4, 3, 90}, target: 0, expected: []int{0, 2}},
		{nums: []int{2, 7}, target: 9, expected: []int{0, 1}},
		{nums: []int{-1_000_000_000, 1_000_000_000}, target: 0, expected: []int{0, 1}},
		{nums: []int{5, 8, 1, 12}, target: 20, expected: []int{1, 3}},
	}

	passed, failed, todo := 0, 0, 0
	fmt.Println("Two Sum practice: implement twoSum(), then run this file.")
	for _, testcase := range cases {
		result := checkOne(testcase.nums, testcase.target)
		input := fmt.Sprintf("nums=%v, target=%d", testcase.nums, testcase.target)
		switch result.status {
		case "PASS":
			passed++
			fmt.Printf("[PASS] %s -> %v\n", input, result.answer)
		case "TODO":
			todo++
			fmt.Printf("[TODO] %s; expected %v (either order)\n", input, testcase.expected)
		case "FAIL":
			failed++
			fmt.Printf("[FAIL] %s -> %v; expected %v (either order)\n", input, result.answer, testcase.expected)
		default:
			failed++
			fmt.Printf("[ERROR] %s -> %s\n", input, result.detail)
		}
	}

	fmt.Printf("Summary: %d PASS, %d FAIL, %d TODO\n", passed, failed, todo)
	if todo > 0 {
		fmt.Println("Replace the TODO in twoSum() with your implementation.")
	}
	if failed > 0 {
		os.Exit(1)
	}
	if todo > 0 {
		os.Exit(2)
	}
}
