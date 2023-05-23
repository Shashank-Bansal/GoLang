package main

func shuffle(nums []int, n int) []int {
	arr := make([]int, 2*n)

	for i, j := 0, 0; i < 2*n; i, j = i+2, j+1 {
		arr[i] = nums[j]
		arr[i+1] = nums[n+j]
	}

	return arr
}
