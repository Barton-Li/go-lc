/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=30204
 *
 * [120] 三角形最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// minimumTotal 计算从三角形顶部到底部的最小路径和。
// triangle: 二维整数数组，表示数字三角形。
// 返回值: 整数，表示从顶部到底部的最小路径和。
func minimumTotal(triangle [][]int) int {
	// 获取三角形的行数
	n := len(triangle)
	// 初始化动态规划表 dp，大小为 n x n
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化起点
	dp[0][0] = triangle[0][0]

	// 填充动态规划表
	for i := 1; i < n; i++ {
		// 处理每行的第一个元素
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		// 处理每行中间的元素
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}

		// 处理每行的最后一个元素
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	// 查找最后一行中的最小值作为结果
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/

