/*
 * @lc app=leetcode.cn id=2658 lang=golang
 * @lcpr version=30204
 *
 * [2658] 网格图中鱼的最大数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y]
		grid[x][y] = 0           // 标记成访问过
		for _, d := range dirs { // 四方向移动
			sum += dfs(x+d.x, y+d.y)
		}
		return sum
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return

}

// @lc code=end

/*
// @lcpr case=start
// [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]\n
// @lcpr case=end

*/

