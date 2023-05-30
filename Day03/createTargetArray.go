package main

func createTargetArray(nums []int, index []int) []int {
	arr := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		arr = insert(arr, index[i], nums[i])
	}

	return arr
}

func insert(arr []int, index int, value int) []int {
	return append(arr[:index], append([]int{value}, arr[index:]...)...)
}
