/*
 * @lc app=leetcode.cn id=1869 lang=golang
 * @lcpr version=30204
 *
 * [1869] 哪种连续子字符串更长
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func checkZeroOnes(s string) bool {
	m1, m0 := 0, 0
	for i, n := 0, len(s); i < n; {
		st := i
		v := s[st]
		for ; i < n && s[i] == v; i++ {
		}
		if v == '1' {
			m1 = max(m1, i-st)
		} else {
			m0 = max(m0, i-st)
		}
	}
	return m1 > m0
}

// @lc code=end

/*
// @lcpr case=start
// "1101"\n
// @lcpr case=end

// @lcpr case=start
// "111000"\n
// @lcpr case=end

// @lcpr case=start
// "110100010"\n
// @lcpr case=end

*/

