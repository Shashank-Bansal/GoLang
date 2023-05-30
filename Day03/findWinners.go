package main

import "sort"

func findWinners(matches [][]int) [][]int {
	set := make(map[int]bool)
	hm := make(map[int]int)

	for _, arr := range matches {
		set[arr[0]] = true
		v := hm[arr[1]]
		hm[arr[1]] = v + 1
	}

	winners := make([]int, 0)
	for i := range set {
		_, ok := hm[i]
		if !ok {
			winners = append(winners, i)
		}
	}

	losers := make([]int, 0)
	for k, v := range hm {
		if v == 1 {
			losers = append(losers, k)
		}
	}

	sort.Ints(winners)
	sort.Ints(losers)

	return [][]int{winners, losers}
}
