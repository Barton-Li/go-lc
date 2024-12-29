/*
 * @lc app=leetcode.cn id=LCR 166 lang=golang
 * @lcpr version=30204
 *
 * [LCR 166] 珠宝的最高价值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// jewelleryValue 计算珠宝框架的最大价值。
// frame 是一个二维整数数组，表示珠宝框架，每个元素代表对应位置的珠宝价值。
// 函数返回从框架的左上角到右下角，按规则移动所能获得的珠宝最大总价值。
func jewelleryValue(frame [][]int) int {
    // 获取框架的行数和列数。
    n, m := len(frame), len(frame[0])
    
    // 初始化动态规划表格，用于记录到每个位置的最大价值。
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }
    
    // 设置起始位置的价值。
    dp[0][0] = frame[0][0]
    
    // 初始化第一列的价值，只能从上往下累加。
    for i := 1; i < n; i++ {
        dp[i][0] = dp[i-1][0] + frame[i][0]
    }
    
    // 初始化第一行的价值，只能从左往右累加。
    for j := 1; j < m; j++ {
        dp[0][j] = dp[0][j-1] + frame[0][j]
    }
    
    // 通过动态规划计算每个位置的最大价值。
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            // 选择从上方或左侧来的最大价值，并加上当前位置的价值。
            dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
        }
    }
    
    // 返回右下角位置的最大价值。
    return dp[n-1][m-1]
}

// max 返回两个整数中的最大值。
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

*/

