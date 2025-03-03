/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	window := [128]bool{}
	left, res := 0, 0
	for right, c := range s {
		for window[c] {
			window[s[left]] = false
			left++
		}
		window[c] = true
		res = max(res, right-left+1)
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/

