package main

func findMaxConsecutiveOnes(nums []int) int {
	n, max := len(nums), 0

	for i := 0; i < n; i++ {
		count := 0

		for i < n && nums[i] == 1 {
			count++
			i++
		}

		if count != 0 && max < count {
			max = count
		}
	}

	return max
}
