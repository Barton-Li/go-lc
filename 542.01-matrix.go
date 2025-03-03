/*
 * @lc app=leetcode.cn id=542 lang=golang
 * @lcpr version=30204
 *
 * [542] 01 矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type point struct {
	x, y int
}

var dirs = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	var res = make([][]int, m)
	var visited = make([][]bool, m)
	var queue []point
	for i := range mat {
		res[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, point{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			x, y := cur.x+dir.x, cur.y+dir.y
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
				res[x][y] = res[cur.x][cur.y] + 1
				queue = append(queue, point{x, y})
				visited[x][y] = true
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[0,1,0],[1,1,1]]\n
// @lcpr case=end

*/

