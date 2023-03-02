package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("ab"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maxValue int
	for i := 0; i < len(s); i++ {
		step := maxValue
		for j := i; j+step <= len(s); j = j + 1 {
			if noRepeat(s[i : j+step]) {
				if j+step-i > maxValue {
					maxValue = j + step - i
				}
			} else {
				break
			}
		}
	}
	return maxValue
}

func noRepeat(subStr string) bool {
	character := map[int32]int{}
	for _, ch := range subStr {
		character[ch] += 1
	}
	return len(character) == len(subStr)
}

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	var l, r, maxLen int
//	var charMap [128]uint8
//
//	for l < len(s) {
//		if r < len(s) && charMap[s[r]] == 0 {
//			charMap[s[r]] = 1
//			r++
//		} else {
//			charMap[s[l]] = 0
//			l++
//		}
//
//		maxLen = max(maxLen, r-l)
//	}
//	return maxLen
//}
//
//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//
//	return b
//}
