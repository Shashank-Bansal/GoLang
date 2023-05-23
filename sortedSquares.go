package main

func sortedSquares(nums []int) []int {
    for i, e := range nums {
        nums[i] = e * e;
    }
    
    return sort(nums);
}

func sort(nums []int) []int {
    n := len(nums);
    
    for i := 0; i < n - 1; i++ {
        for j := 0; j < n - 1 - i; j++ {
            if (nums[j] > nums[j + 1]) {
                nums[j], nums[j + 1] = nums[j + 1], nums[j];
            }
        }
    }
    
    return nums;
}