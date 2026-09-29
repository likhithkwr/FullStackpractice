package main

import (
	"slices"
	"testing"
)

func requireValidPair(t *testing.T, nums []int, target int, answer []int) {
	t.Helper()
	if len(answer) != 2 {
		t.Fatalf("nums=%v, target=%d: expected two indices, got %v", nums, target, answer)
	}

	first, second := answer[0], answer[1]
	if first < 0 || first >= len(nums) || second < 0 || second >= len(nums) || first == second {
		t.Fatalf("nums=%v, target=%d: invalid indices %v", nums, target, answer)
	}
	if nums[first]+nums[second] != target {
		t.Fatalf("nums=%v, target=%d: indices %v do not sum to the target", nums, target, answer)
	}
}

func TestTwoSumApproaches(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
	}{
		{name: "first example", nums: []int{2, 7, 11, 15}, target: 9},
		{name: "middle pair", nums: []int{3, 2, 4}, target: 6},
		{name: "duplicate values", nums: []int{3, 3}, target: 6},
		{name: "zero values", nums: []int{0, 4, 3, 0}, target: 0},
		{name: "negative values", nums: []int{-3, 4, 3, 90}, target: 0},
		{name: "two elements", nums: []int{2, 7}, target: 9},
		{name: "input bounds", nums: []int{-1_000_000_000, 1_000_000_000}, target: 0},
		{name: "last pair", nums: []int{5, 8, 1, 12}, target: 20},
	}
	approaches := []struct {
		name  string
		solve func([]int, int) []int
	}{
		{name: "hash map", solve: twoSum},
		{name: "brute force", solve: twoSumBruteForce},
	}

	for _, approach := range approaches {
		t.Run(approach.name, func(t *testing.T) {
			for _, testcase := range cases {
				t.Run(testcase.name, func(t *testing.T) {
					input := slices.Clone(testcase.nums)
					answer := approach.solve(input, testcase.target)
					if !slices.Equal(input, testcase.nums) {
						t.Fatalf("%s changed the input array", approach.name)
					}
					requireValidPair(t, testcase.nums, testcase.target, answer)
				})
			}
		})
	}
}
