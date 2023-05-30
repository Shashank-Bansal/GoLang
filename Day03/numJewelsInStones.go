package main

func numJewelsInStones(jewels string, stones string) int {
	hms := createMap(stones)

	count := 0
	for _, e := range jewels {
		count += hms[e]
	}

	return count
}

func createMap(s string) map[rune]int {
	hm := make(map[rune]int)

	for _, e := range s {
		v := hm[e]
		hm[e] = v + 1
	}

	return hm
}
