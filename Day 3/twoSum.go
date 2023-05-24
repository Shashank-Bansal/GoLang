package main

// // method 1
// func twoSum(nums []int, target int) []int {
//     for i := 0; i < len(nums) - 1; i++ {
//         for j := i + 1; j < len(nums); j++ {
//             if nums[i] + nums[j] == target {
//                 return []int{i, j};
//             }
//         }
//     }

//     return []int{-1, -1};
// }

// // method 2
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, e := range nums {
		v, ok := hashmap[e]
		if ok {
			return []int{v, i}
		}

		hashmap[target-e] = i
	}

	return []int{-1, -1}
}
