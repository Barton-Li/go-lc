/*
 * @lc app=leetcode.cn id=740 lang=golang
 * @lcpr version=30204
 *
 * [740] 删除并获得点数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// deleteAndEarn 函数接收一个整数数组 nums，返回通过删除某些元素以最大化得分的方式所得到的最大得分。
// 这个函数的核心思想是通过动态规划来解决问题，其中涉及到对数组元素的排序和分情况讨论以决定如何最大化得分。
func deleteAndEarn(nums []int) (ans int) {
	// 对数组 nums 进行排序，以便后续处理。
	sort.Ints(nums)

	// 初始化 sum 切片用于记录连续元素的和。
	sum := []int{nums[0]}

	// 遍历排序后的数组 nums。
	for i := 1; i < len(nums); i++ {
		// 使用 var 变量简化代码，并提高可读性。
		if val := nums[i]; val == nums[i-1] {
			// 如果当前元素与前一个元素相同，将其值加到 sum 的最后一个元素上。
			sum[len(sum)-1] += val
		} else if val == nums[i-1]+1 {
			// 如果当前元素是前一个元素的值加一，将当前值作为一个新的和添加到 sum 中。
			sum = append(sum, val)
		} else {
			// 如果当前元素既不等于前一个元素，也不等于前一个元素加一，
			// 则将 sum 分割为两部分，并计算第一部分的最大得分。
			ans += rob(sum)
			// 重新开始一个新的 sum 切片，只包含当前元素值。
			sum = []int{val}
		}
	}
	// 在遍历结束后，计算最后一部分 sum 的最大得分，并将其加到 ans 中。
	ans += rob(sum)
	return
}
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
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
// [3,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,3,3,3,4]\n
// @lcpr case=end

*/

