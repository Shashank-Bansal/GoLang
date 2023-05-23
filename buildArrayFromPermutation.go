package main

func buildArray(nums []int) []int {
    n := len(nums);
    arr := make([]int, n);
    
    for i, _ := range nums {
        arr[i] = nums[nums[i]];
    }
    
    return arr;
}