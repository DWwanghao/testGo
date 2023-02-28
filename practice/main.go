package main

import (
	"fmt"
)

func main() {
	//fmt.Println(lengthOfLongestSubstring("bbcca"))
	fmt.Println(len("a6z"))
	fmt.Println(lengthOfLongestSubstring("a6z"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maxValue int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			temp := s[i:j]
			fmt.Println(temp)
			if noRepeat(s[i:j]) {
				if j-i > maxValue {
					maxValue = j - i
				}
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
	for _, i := range character {
		if i > 1 {
			return false
		}
	}
	return true

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
