package main

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	sum, i := 0, 0

	for i < k {
		sum += nums[i]
		i++
	}

	var max = float64(sum) / float64(k)

	for ; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		max = math.Max(max, float64(sum)/float64(k))
	}

	return max
}
