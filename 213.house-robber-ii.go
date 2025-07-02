/*
 * @lc app=leetcode.cn id=213 lang=golang
 * @lcpr version=30204
 *
 * [213] 打家劫舍 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robx(nums[:n-1]), robx(nums[1:]))

}
// robx 函数计算在不触发警报的情况下，从一排房屋中能抢劫到的最大金额。
// 它接受一个整数数组 nums 作为输入，数组中的每个元素代表一所房屋内的金额。
// 函数返回能抢劫到的最大金额。
func robx(nums []int) int {
    // 初始化 first 为第一所房屋的金额，second 为前两所房屋中金额较大的一个。
    first, second := nums[0], max(nums[0], nums[1])

    // 遍历房屋列表，从第三所房屋开始。
    // 使用范围循环遍历 nums[2:]，v 代表当前房屋的金额。
    for _, v := range nums[2:] {
        // 更新 first 和 second。
        // first 变为上一个 second，second 变为抢劫当前房屋和上一个房屋的最大金额。
        first, second = second, max(first+v, second)
    }

    // 返回能抢劫到的最大金额。
    return second
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

