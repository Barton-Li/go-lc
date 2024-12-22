/*
 * @lc app=leetcode.cn id=1191 lang=golang
 * @lcpr version=30204
 *
 * [1191] K 次串联后最大子数组之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
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
func kConcatenationMaxSum(arr []int, k int) int {
	const mod = 1_000_000_007
	//把arr数组重复k次后 求最大的子数组和
	//假设2个arr数组的最大子数组和为maxSum，arr数组的元素和为sum，那么最后的答案为max(sum,0) * (k - 2) + maxSum
	sum := 0
	for _, num := range arr {
		sum = (sum + num) % mod
	}
	if k >= 2 {
		arr = append(arr, arr...)
	}
	n, maxSum := len(arr), 0
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = max(dp[i-1]+arr[i-1], arr[i-1])
		maxSum = max(maxSum, dp[i])
	}
	if k <= 2 {
		return maxSum % mod
	}
	return (max(sum, 0)*(k-2)%mod + maxSum%mod) % mod
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

