/*
 * @lc app=leetcode.cn id=1254 lang=golang
 * @lcpr version=30204
 *
 * [1254] 统计封闭岛屿的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// closedIsland 计算封闭岛屿的数量。
// grid 是一个二维数组，其中0表示陆地，1表示水域。
// 封闭岛屿是指完全由水域包围的陆地。
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0]) // m和n分别代表网格的行数和列数

	// dfs 是深度优先搜索函数，用于标记从(i, j)开始的连通陆地。
	// 如果(i, j)超出网格边界或已经是水域，则返回。
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 0 {
			return
		}
		grid[i][j] = 1 // 将陆地标记为水域，避免重复计数
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	// 首先，从网格的边界开始，标记所有与边界相连的陆地。
	// 这样做是为了排除那些非封闭的陆地。
	// 遍历矩阵的每个行
	for i := 0; i < m; i++ {
		// 初始化步长为1
		step := 1
		// 如果不是第一行或最后一行，步长设置为n-1
		if 0 < i && i < m-1 {
			step = n - 1
		}
		// 在当前行中，以步长step遍历每个列
		for j := 0; j < n; j += step {
			// 调用dfs函数处理当前位置
			dfs(i, j)
		}
	}

	// 然后，遍历网格内部，寻找未被标记的封闭岛屿。
	// 每找到一个封闭岛屿，就增加计数并标记整个岛屿。
	res := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1,1,1,1,1],[1,0,0,0,0,0,1],[1,0,1,1,1,0,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,0,0,0,0,0,1],[1,1,1,1,1,1,1]]\n
// @lcpr case=end

*/

