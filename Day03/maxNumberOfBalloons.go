package main

import (
	"math"
	"strings"
)

func maxNumberOfBalloons(text string) int {
	s := "balloon"
	hms := createMap(s)
	hm := createMap(text)

	min := math.MaxInt

	for _, e := range s {
		m := hm[e] / hms[e]
		if m < min {
			min = m
		}
	}
	
	return min
	
	// if hm['l'] == min*2 && hm['o'] == min*2 {
	// 	return min
	// }

	// secondMin := minValue(hm['l'], hm['o'])

	// if secondMin%2 != 0 {
	// 	return secondMin - 1
	// }

	// return secondMin
}

func createMap(s string) map[rune]int {
	hm := make(map[rune]int)

	for _, e := range s {
		if strings.ContainsRune(s, e) {
			v := hm[e]
			hm[e] = v + 1
		}
	}

	return hm
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
