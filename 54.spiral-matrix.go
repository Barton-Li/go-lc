/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30204
 *
 * [54] 螺旋矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int,  m*n)
	i, j, di := 0, 0, 0
	for k := range res {
		res[k] = matrix[i][j]
		matrix[i][j] = math.MaxInt
		x, y := i+dirs[di][0], j+dirs[di][1]
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] == math.MaxInt {
			di = (di + 1) % 4
		}
		i += dirs[di][0]
		j += dirs[di][1]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/

