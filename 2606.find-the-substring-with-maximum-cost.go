/*
 * @lc app=leetcode.cn id=2606 lang=golang
 * @lcpr version=30204
 *
 * [2606] 找到最大开销的子字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumCostSubstring(s string, chars string, vals []int) int {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	var res int
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		res = max(res, f)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "adaa"\n"d"\n[-1000]\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n[-1,-1,-1]\n
// @lcpr case=end

*/

