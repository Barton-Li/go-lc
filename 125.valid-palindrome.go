/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30204
 *
 * [125] 验证回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if right > left {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isalnum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/

