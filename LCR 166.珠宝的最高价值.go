/*
 * @lc app=leetcode.cn id=LCR 166 lang=golang
 * @lcpr version=30204
 *
 * [LCR 166] 珠宝的最高价值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func jewelleryValue(frame [][]int) int {
	n, m := len(frame), len(frame[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = frame[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + frame[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + frame[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
		}
	}
	return dp[n-1][m-1]

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

*/

