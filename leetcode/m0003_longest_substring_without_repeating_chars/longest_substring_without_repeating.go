package longestsubstringwithoutrepeatingchars

import "practice/utility/math"

// Time complexity: O(N)
// Space complexity: O(N)
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	ans := 0
	count := make(map[byte]int)
	for right < len(s) {
		count[s[right]]++
		for left < right && count[s[right]] > 1 {
			count[s[left]]--
			left++
		}

		ans = math.Max(ans, right-left+1)

		right++
	}

	return ans
}

// Time complexity: O(N)
// Space complexity: O(N)
func lengthOfLongestSubstringOptimized(s string) int {
	ans := 0
	count := make(map[byte]int)
	for left, right := 0, 0; right < len(s); right++ {
		if index, exist := count[s[right]]; exist && index >= left {
			left = index + 1
		}

		ans = math.Max(ans, right-left+1)
		count[s[right]] = right
	}

	return ans
}
