/*
 * @lc app=leetcode.cn id=1446 lang=golang
 * @lcpr version=30204
 *
 * [1446] 连续字符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxEnergy := 0
	currentEnergy := 0
	prevChar := rune(0)
	for _, char := range s {
		if char == prevChar {
			currentEnergy++
		} else {
			maxEnergy = max(maxEnergy, currentEnergy)
			currentEnergy = 1
			prevChar = char

		}
	}
	maxEnergy = max(maxEnergy, currentEnergy)
	return maxEnergy
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

// @lcpr case=start
// "abbcccddddeeeeedcba"\n
// @lcpr case=end

*/

