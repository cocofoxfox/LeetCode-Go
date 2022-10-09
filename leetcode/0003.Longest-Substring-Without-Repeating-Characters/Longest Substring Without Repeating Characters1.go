package leetcode

// memorized array
func lengthOfLongestSubstring4(s string) int {
	var seen [256]bool
	left, right, max := 0, 0, 0
	for left < len(s) && right < len(s) && max <= len(s)-left {
		if seen[s[right]] {
			seen[s[left]] = false
			left++
		} else {
			seen[s[right]] = true
			right++
			if max < right-left {
				max = right - left
			}
		}
	}
	return max
}

// sliding window
func lengthOfLongestSubstring5(s string) int {
	if len(s) == 0 {
		return 0
	}
	lookup := make(map[byte]int)
	left, max := 0, 0

	for right := range s {
		if _, exists := lookup[s[right]]; exists && left <= lookup[s[right]] {
			left = lookup[s[right]] + 1
		}
		if max < right-left+1 {
			max = right - left + 1
		}
		lookup[s[right]] = right
	}
	return max
}
