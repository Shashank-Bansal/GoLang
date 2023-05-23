package main

func runningSum(nums []int) []int {
	var sum int
	for i, e := range nums {
		sum += e
		nums[i] = sum
	}

	return nums
}
