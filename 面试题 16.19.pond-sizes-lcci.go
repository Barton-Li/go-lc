/*
 * @lc app=leetcode.cn id=面试题 16.19 lang=golang
 * @lcpr version=30204
 *
 * [面试题 16.19] 水域大小
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func pondSizes(land [][]int) []int {
	m, n := len(land), len(land[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || land[i][j] != 0 {
			return 0
		}
		land[i][j] = 1
		cnt := 1
		for x := i - 1; x <= i+1; x++ {
			for y := j - 1; y <= j+1; y++ {
				cnt += dfs(x, y)
			}

		}
		return cnt
	}
	res := make([]int, 0)
	for i, v := range land {
		for j, x := range v {
			if x == 0 {
				res = append(res, dfs(i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[0,2,1,0],[0,1,0,1],[1,1,0,1],[0,1,0,1]]\n
// @lcpr case=end

*/

