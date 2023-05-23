package main
 
func removeElement(nums []int, val int) int {
    i := len(nums) - 1;
    
    for j := 0; j <= i; j++ {
        if (nums[j] == val) {
            nums[j] = nums[i];
            j--;
            i--;
        }
    }
    
    return i + 1;
}