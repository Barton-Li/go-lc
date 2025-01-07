/*
 * @lc app=leetcode.cn id=463 lang=golang
 * @lcpr version=30204
 *
 * [463] 岛屿的周长
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func islandPerimeter(grid [][]int) int {
	per := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i == 0 || grid[i-1][j] == 0 {
					per++
				}
				if i == m-1 || grid[i+1][j] == 0 {
					per++
				}
				if j == 0 || grid[i][j-1] == 0 {
					per++
				}
				if j == n-1 || grid[i][j+1] == 0 {
					per++
				}
			}
		}
	}
	return per
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0]]\n
// @lcpr case=end

*/

