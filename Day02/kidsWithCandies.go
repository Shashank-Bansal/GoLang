package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := greatest(candies)
	n := len(candies)
	arr := make([]bool, n)
	
	for i, e := range candies {
		if e+extraCandies >= max {
			arr[i] = true
		}
	}

	return arr
}

func greatest(arr []int) int {
	var max int

	for _, e := range arr {
		if max < e {
			max = e
		}
	}

	return max
}
