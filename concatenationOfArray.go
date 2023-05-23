package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	arr := make([]int, 2*n)

	for i, e := range nums {
		arr[i] = e
		arr[i+n] = e
	}

	return arr
}
