/*
 * @lc app=leetcode.cn id=377 lang=golang
 * @lcpr version=30204
 *
 * [377] 组合总和 Ⅳ
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func combinationSum4(nums []int, target int) int {
// 	f := make([]int, target+1)
// 	f[0] = 1
// 	for i := 1; i <= target; i++ {
// 		for _, x := range nums {
// 			if x <= i {
// 				f[i] += f[i-x]
// 			}
// 		}
// 	}
// 	return f[target]
// }
// func combinationSum4(nums []int, target int) int {
// 	memu := make([]int, target+1)
// 	for i := range memu {
// 		memu[i] = -1
// 	}
// 	var dfs func(int) int
// 	dfs = func(i int) (res int) {
// 		if i == 0 {
// 			return 1
// 		}
// 		p := memu[i]
// 		if p != -1 {
// 			return p
// 		}
// 		for _, x := range nums {
// 			if x <= i {
// 				res += dfs(i - x)
// 			}
// 		}
// 		p = res
// 		return
// 	}
// 	return dfs(target)
// }
// combinationSum4 calculates the number of combinations in the given array nums that sum up to the target.
// This function uses a dynamic programming approach to solve the problem, avoiding repeated calculations and backtracking.
// Parameters:
//   nums: An integer array containing a set of candidate numbers.
//   target: The target sum, the function aims to find the number of combinations in nums that sum up to the target.
// Return value:
//   Returns an integer, representing the number of combinations that sum up to the target.
func combinationSum4(nums []int, target int) int {
    // Initialize the dp array with a length of target+1, to accommodate all possible sums from 0 to target.
    dp := make([]int, target+1)
    // dp[0] = 1 is the base case, indicating that there is 1 combination for a sum of 0 (i.e., choosing no elements).
    dp[0] = 1

    // The outer loop iterates through all possible sums from 1 to target.
    for i := 1; i <= target; i++ {
        // The inner loop iterates through each element in the nums array.
        for _, num := range nums {
            // If the current sum i is greater than or equal to the element num in nums, num can be part of the combination.
            if i >= num {
                // Update dp[i], adding the number of combinations ending with num.
                dp[i] += dp[i-num]
            }
        }
    }
    // Return the number of combinations for the target sum.
    return dp[target]
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [9]\n3\n
// @lcpr case=end

*/

