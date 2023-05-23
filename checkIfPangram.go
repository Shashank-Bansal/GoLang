package main

// // method 1
// func checkIfPangram(sentence string) bool {
//     arr := make([] bool, 26)

//     for _, e := range sentence {
//         arr[e - 'a'] = true;
//     }

//     for _, e := range arr {
//         if !e {
//             return false;
//         }
//     }

//     return true;
// }

// // method 2

import (
	"strings"
)

func checkIfPangram(sentence string) bool {
	for i := 0; i < 26; i++ {
		if !strings.Contains(sentence, string(rune('a'+i))) {
			return false
		}
	}

	return true
}
