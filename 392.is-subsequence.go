/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=30204
 *
 * [392] 判断子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/

