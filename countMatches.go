package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int

	for _, e := range items {
		if (ruleKey == "type" && ruleValue == e[0]) || (ruleKey == "color" && ruleValue == e[1]) || (ruleKey == "name" && ruleValue == e[2]) {
			count++
		}
	}

	return count
}
