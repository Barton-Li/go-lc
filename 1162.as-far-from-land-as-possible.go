/*
 * @lc app=leetcode.cn id=1162 lang=golang
 * @lcpr version=30204
 *
 * [1162] 地图分析
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func maxDistance(grid [][]int) int {
	n := len(grid)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1

		}
	}
	queue := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
				dist[i][j] = 0
			}
		}
	}
	if len(queue) == n*n {
		reuturn - 1
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && dist[nx][ny] == -1 {
				dist[nx][ny] = dist[x][y] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}
	maxDist := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dist[i][j] > maxDist {
				maxDist = dist[i][j]
			}
		}
	}
	return maxDist
}

// @lc code=end

/*
// @lcpr case=start
// [[1,0,1],[0,0,0],[1,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[0,0,0],[0,0,0]]\n
// @lcpr case=end

*/

