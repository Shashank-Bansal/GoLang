package main

func lengthOfLongestSubstring(s string) int {
	set := []rune{}
	var max int

	for _, char := range s {
		prevIndex := contains(set, char)
		if prevIndex == -1 {
			set = append(set, char)
		} else {
			if len(set) > max {
				max = len(set)
			}

			set = set[prevIndex+1:]
			set = append(set, char)
		}
	}

	if len(set) > max {
		max = len(set)
	}

	return max
}

func contains(set []rune, char rune) int {
	for index, element := range set {
		if element == char {
			return index
		}
	}

	return -1
}
