/*
 * @lc app=leetcode.cn id=62 lang=golang
 * @lcpr version=30204
 *
 * [62] 不同路径
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// uniquePaths 计算从矩阵的左上角到右下角的独特路径数量。
// 参数：
//   m - 矩阵的行数
//   n - 矩阵的列数
// 返回值：
//   从左上角到右下角的独特路径数量
func uniquePaths(m int, n int) int {
    // 初始化一个二维切片 dp，用于存储到达每个位置的独特路径数量
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    for i := 0; i < n; i++ {
        dp[0][i] = 1
    }
    // 从 (1,1) 开始，计算到达每个位置的独特路径数量
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 当前位置的独特路径数量等于上方和左方路径数量之和
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    // 返回到达右下角的独特路径数量
    return dp[m-1][n-1]
}
// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n
// @lcpr case=end

*/

