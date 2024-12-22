/*
 * @lc app=leetcode.cn id=2321 lang=golang
 * @lcpr version=30204
 *
 * [2321] 拼接数组的最大分数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// maximumsSplicedArray 计算两个数组通过交换子数组后可以获得的最大和。
// 参数：
//   nums1：第一个数组
//   nums2：第二个数组
// 返回值：
//   交换子数组后可以获得的最大和
func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	// 分别计算将 nums1 和 nums2 中的某个子数组进行交换后的最大和，
	// 取两者中的较大值作为最终结果。
	return max(maxSwap(nums1, nums2), maxSwap(nums2, nums1))
}

// maxSwap 计算在数组 nums1 中选择一个子数组与 nums2 对应位置的子数组交换后，
// 能够获得的最大和增量。
// 参数：
//   nums1：原始数组
//   nums2：用于交换的数组
// 返回值：
//   交换子数组后，nums1 的最大和增量
func maxSwap(nums1 []int, nums2 []int) int {
	var sum, maxSum, f int
	// 遍历数组，计算 nums1 的总和，并计算通过交换子数组可以获得的最大增量。
	for i, num := range nums1 {
		sum += num
		// f 表示当前子数组交换后的增量，使用动态规划思想更新 f 和 maxSum。
		f = max(f, 0) + nums2[i] - num
		maxSum = max(maxSum, f)
	}
	// 返回 nums1 的总和加上最大增量。
	return sum + maxSum
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [60,60,60]\n[10,90,10]\n
// @lcpr case=end

// @lcpr case=start
// [20,40,20,70,30]\n[50,20,50,40,20]\n
// @lcpr case=end

// @lcpr case=start
// [7,11,13]\n[1,1,1]\n
// @lcpr case=end

*/

