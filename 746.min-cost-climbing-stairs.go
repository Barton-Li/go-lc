/*
 * @lc app=leetcode.cn id=746 lang=golang
 * @lcpr version=30204
 *
 * [746] 使用最小花费爬楼梯
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minCostClimbingStairs(cost []int) int {
	f0, f1 := 0, 0
	for i := 1; i < len(cost); i++ {
		f0, f1 = f1, min(f1+cost[i], f0+cost[i-1])
	}
	return f1
}

// @lc code=end

/*
// @lcpr case=start
// [10,15,20]\n
// @lcpr case=end

// @lcpr case=start
// [1,100,1,1,1,100,1,1,100,1]\n
// @lcpr case=end

*/

