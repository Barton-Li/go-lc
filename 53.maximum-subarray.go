/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30204
 *
 * [53] 最大子数组和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]+nums[i-1] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}

// 	}
// 	return max
// }
func maxSubArray(nums []int) int {
	res := math.MinInt32
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		res = max(res, f)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

