/*
 * @lc app=leetcode.cn id=1749 lang=golang
 * @lcpr version=30204
 *
 * [1749] 任意子数组和的绝对值的最大值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxAbsoluteSum(nums []int) int {
	var res int
	fMin, fMax := 0, 0
	for _, num := range nums {
		fMax = max(fMax, 0) + num
		fMin = min(fMin, 0) + num
		res = max(res, max(fMax, -fMin))
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,-3,2,3,-4]\n
// @lcpr case=end

// @lcpr case=start
// [2,-5,1,-4,3,-2]\n
// @lcpr case=end

*/

