/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=30204
 *
 * [63] 不同路径 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// uniquePathsWithObstacles 计算从左上角到右下角的无障碍路径数量。
// obstacleGrid 是一个二维数组，其中 0 代表无障碍，1 代表有障碍。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // n, m 分别代表网格的行数和列数。
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    // dp 数组用于动态规划，dp[j] 表示到达第 j 列的路径数。
    dp := make([]int, m)
    // 如果起始位置没有障碍，则初始化 dp[0] 为 1。
    if obstacleGrid[0][0] == 0 {
        dp[0] = 1
    }
    // 遍历网格的每一行。
    for i := 0; i < n; i++ {
        // 遍历网格的每一列。qa
        for j := 0; j < m; j++ {
            // 如果当前位置有障碍，则到达该位置的路径数为 0，并继续下一次循环。
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
                continue
            }
            // 如果当前位置的左侧没有障碍，则当前位置的路径数包括从左侧过来的路径数。
            if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
                dp[j] += dp[j-1]
            }
        }
    }
    // 返回到达右下角的路径数量。
    return dp[m-1]
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

*/

