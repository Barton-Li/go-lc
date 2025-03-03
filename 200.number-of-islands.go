/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30204
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// numIslands 计算网格中岛屿的数量。
// grid 是一个二维网格，其中 '1' 代表陆地，'0' 代表水域。
// 岛屿是通过水平或垂直连接的陆地形成的。
func numIslands(grid [][]byte) int {
	// m 和 n 分别是网格的行数和列数。
	m, n := len(grid), len(grid[0])

	// dfs 是一个递归函数，用于深度优先搜索岛屿的边界。
	// 它通过将访问过的陆地标记为 '2' 来避免重复计算。
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 如果当前位置超出网格边界或不是陆地，则返回。
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		// 将当前位置标记为已访问。
		grid[i][j] = '2'
		// 递归地访问所有相邻的陆地。
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}

	// sum 用于记录岛屿的数量。
	sum := 0
	// 遍历网格中的每个位置。
	for i, row := range grid {
		for j, v := range row {
			// 如果当前位置是未访问的陆地，则进行深度优先搜索。
			if v == '1' {
				dfs(i, j)
				// 搜索完成后，岛屿数量加一。
				sum++
			}
		}
	}
	// 返回岛屿的数量。
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/

