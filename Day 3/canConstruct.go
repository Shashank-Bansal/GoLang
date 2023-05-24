package main

func canConstruct(ransomNote string, magazine string) bool {
	hmr := createMap(ransomNote)
	hmm := createMap(magazine)

	for k, v := range hmr {
		if v < hmm[k] || hmm[k] == 0 {
			return false
		}
	}

	return true
}

func createMap(s string) map[rune]int {
	hm := make(map[rune]int)

	for _, e := range s {
		v := hm[e]
		hm[e] = v + 1
	}

	return hm
}
