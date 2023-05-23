package main

func smallerNumbersThanCurrent(nums []int) []int {
	var n int = len(nums)
	var arr []int = make([]int, n)

	for i, e := range nums {
		arr[i] = counter(nums, e)
	}

	return arr
}

func counter(nums []int, t int) int {
	var count int

	for _, e := range nums {
		if e < t {
			count++
		}
	}

	return count
}
