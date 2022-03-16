package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcaabccc"))
}

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
