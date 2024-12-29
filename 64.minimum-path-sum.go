/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=30204
 *
 * [64] 最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minPathSum(grid [][]int) int {
	// 获取网格的行数和列数
	n, m := len(grid), len(grid[0])

	// 初始化动态规划表 dp，大小与 grid 相同
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 设置起点的值
	dp[0][0] = grid[0][0]

	// 初始化第一列的值：只能从上方到达
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 初始化第一行的值：只能从左侧到达
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 计算其余位置的最小路径和：可以从上方或左侧到达
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	// 返回到达右下角的最小路径和
	return dp[n-1][m-1]
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/

