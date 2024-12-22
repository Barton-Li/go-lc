/*
 * @lc app=leetcode.cn id=2266 lang=golang
 * @lcpr version=30204
 *
 * [2266] 统计打字方案数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countTexts(pressedKeys string) int {
	mod := 1000000007
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = (dp[i] + dp[i-1]) % mod
		if i >= 2 && pressedKeys[i-1] == pressedKeys[i-2] {
			dp[i] = (dp[i] + dp[i-2]) % mod
		}
		if i >= 3 && pressedKeys[i-1] == pressedKeys[i-2] && pressedKeys[i-1] == pressedKeys[i-3] {
			dp[i] = (dp[i] + dp[i-3]) % mod
		}
		if i >= 4 && (pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9') &&
			pressedKeys[i-1] == pressedKeys[i-2] && pressedKeys[i-1] == pressedKeys[i-3] &&
			pressedKeys[i-1] == pressedKeys[i-4] {
			dp[i] = (dp[i] + dp[i-4]) % mod
		}
	}

	return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// "22233"\n
// @lcpr case=end

// @lcpr case=start
// "222222222222222222222222222222222222"\n
// @lcpr case=end

*/

