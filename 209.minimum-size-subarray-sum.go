/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30204
 *
 * [209] 长度最小的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

