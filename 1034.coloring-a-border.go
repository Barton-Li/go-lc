/*
 * @lc app=leetcode.cn id=1034 lang=golang
 * @lcpr version=30204
 *
 * [1034] 边界着色
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row, col, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	borders := []point{}
	originalColor := grid[row][col]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isBorder := false
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBorder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)

	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1],[1,2]]\n0\n0\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,2],[2,3,2]]\n0\n1\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1],[1,1,1],[1,1,1]]\n1\n1\n2\n
// @lcpr case=end

*/

