/*
 * @lc app=leetcode.cn id=2466 lang=golang
 * @lcpr version=30204
 *
 * [2466] 统计构造好字符串的方案数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000007
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n3\n1\n1\n
// @lcpr case=end

// @lcpr case=start
// 2\n3\n1\n2\n
// @lcpr case=end

*/

