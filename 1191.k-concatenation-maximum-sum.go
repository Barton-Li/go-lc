/*
 * @lc app=leetcode.cn id=1191 lang=golang
 * @lcpr version=30204
 *
 * [1191] K 次串联后最大子数组之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}

// kConcatenationMaxSum 计算将数组 arr 重复 k 次后，最大子数组的和。
// 参数：
//   arr：原始数组
//   k：数组重复的次数
// 返回值：
//   最大子数组的和，结果取模 1_000_000_007
func kConcatenationMaxSum(arr []int, k int) int {
	// 定义常量 mod 用于对结果取模
	const mod = 1_000_000_007

	// 计算数组 arr 的元素和，并对 mod 取模
	sum := 0
	for _, num := range arr {
		sum = (sum + num) % mod
	}

	// 如果 k 大于等于 2，则将 arr 数组与其自身连接
	if k >= 2 {
		arr = append(arr, arr...)
	}

	// 初始化数组长度 n 和最大子数组和 maxSum
	maxSum, f := 0, 0

	for _, num := range arr {
		f = max(f, 0) + num
		maxSum = max(maxSum, f)
	}

	// 如果 k 不大于 2，直接返回最大子数组和并对 mod 取模
	if k <= 2 {
		return maxSum % mod
	}

	// 根据 sum 和 maxSum 计算最终结果，并对 mod 取模
	return (max(sum, 0)*(k-2)%mod + maxSum%mod) % mod
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
// [1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,-2,1]\n5\n
// @lcpr case=end

// @lcpr case=start
// [-1,-2]\n7\n
// @lcpr case=end

*/

