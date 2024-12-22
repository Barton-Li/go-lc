/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=30204
 *
 * [152] 乘积最大子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProduct(nums []int) int {
	res := math.MinInt32
	fMax, fMin := 1, 1
	for _, num := range nums {
		fMax, fMin = max(fMax*num, fMin*num, num), min(fMax*num, fMin*num, num)
		res = max(res, fMax)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/

