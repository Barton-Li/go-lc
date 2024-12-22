/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=30204
 *
 * [198] 打家劫舍
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// rob 函数计算盗窃一列房屋中的最大金额
// 参数 nums 是一个整数数组，表示每间房屋内的金额
// 返回值是盗窃房屋能得到的最大金额
func rob(nums []int) int {
	// 如果没有房屋，返回0
	if len(nums) == 0 {
		return 0
	}
	// 如果只有一间房屋，返回该房屋的金额
	if len(nums) == 1 {
		return nums[0]
	}
	// 初始化动态规划数组dp，长度与房屋数量相同
	dp := make([]int, len(nums))
	// 第一间房屋的最大金额就是其自身的金额
	dp[0] = nums[0]
	// 第二间房屋的最大金额是前两间房屋中金额较大的一间
	dp[1] = max(nums[0], nums[1])
	// 从第三间房屋开始，使用动态规划计算每间房屋的最大金额
	for i := 2; i < len(nums); i++ {
		// 对于当前房屋，有两种选择：盗窃或不盗窃
		// 如果盗窃当前房屋，则前一间房屋不能盗窃，因此最大金额是前两间房屋的最大金额加上当前房屋的金额
		// 如果不盗窃当前房屋，则最大金额是前一间房屋的最大金额
		// 选择这两种情况中的最大值作为当前房屋的最大金额
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	// 返回最后一间房屋的最大金额作为最终结果
	return dp[len(nums)-1]
}

// max 函数返回两个整数中的较大值
// 参数 a 和 b 是两个整数
// 返回值是 a 和 b 中的较大值
func max(a, b int) int {
	// 如果 a 大于 b，返回 a
	if a > b {
		return a
	}
	// 否则，返回 b
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/

