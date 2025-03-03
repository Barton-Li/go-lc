/*
 * @lc app=leetcode.cn id=205 lang=golang
 * @lcpr version=30204
 *
 * [205] 同构字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "egg"\n"add"\n
// @lcpr case=end

// @lcpr case=start
// "foo"\n"bar"\n
// @lcpr case=end

// @lcpr case=start
// "paper"\n"title"\n
// @lcpr case=end

*/

