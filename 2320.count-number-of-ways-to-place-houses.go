/*
 * @lc app=leetcode.cn id=2320 lang=golang
 * @lcpr version=30204
 *
 * [2320] 统计放置房子的方式数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func countHousePlacements(n int) int {
	const mod = 1e9 + 7
	// f(n) = f(n-1) + f(n-2) m = p + q
	p, q, m := 2, 1, 3
	if n == 1 {
		m = 2
	}
	for i := 2; i <= n; i++ {
		m = (p + q) % mod
		q = p
		p = m
	}
	return m * m % mod
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/

