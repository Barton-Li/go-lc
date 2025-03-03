/*
 * @lc app=leetcode.cn id=695 lang=golang
 * @lcpr version=30204
 *
 * [695] 岛屿的最大面积
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// maxAreaOfIsland 计算给定二维网格中最大岛屿的面积。
// grid 二维整数数组，表示网格，其中1表示陆地，0表示水域。
// 返回最大岛屿的面积。
func maxAreaOfIsland(grid [][]int) int {
    // 获取网格的行数和列数
    m, n := len(grid), len(grid[0])

    // dfs 是一个递归函数，用于深度优先搜索岛屿的面积。
    // i, j 表示当前网格的位置。
    // 返回从当前位置开始的岛屿部分的面积。
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        // 检查当前位置是否越界或为水域，如果是，则返回0。
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
            return 0
        }
        // 将当前位置标记为已访问，避免重复计算。
        grid[i][j] = 2
        // 递归计算上下左右四个方向的岛屿面积，并加上当前位置的面积1。
        return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
    }

    // 初始化最大面积为0。
    var res int
    // 遍历网格的每个位置。
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // 更新最大岛屿面积。
            res = max(res, dfs(i, j))
        }
    }
    // 返回最大岛屿面积。
    return res
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
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0,0,0,0,0,0]]\n
// @lcpr case=end

*/

