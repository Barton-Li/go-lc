/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=30204
 *
 * [242] 有效的字母异位词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isAnagram(s string, t string) bool {
	sres, tres := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		sres[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		tres[t[i]-'a']++
	}
	return sres == tres
}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/

