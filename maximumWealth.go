package main

func maximumWealth(accounts [][]int) int {
	var max int

	for _, e := range accounts {
		s := sum(e)
		if max < s {
			max = s
		}
	}

	return max
}

func sum(arr []int) int {
	var sum int

	for _, e := range arr {
		sum += e
	}

	return sum
}
