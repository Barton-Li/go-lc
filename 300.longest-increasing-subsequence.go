/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30204
 *
 * [300] 最长递增子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func lengthOfLIS2(nums []int)int{
	n:=len(nums)
	if n:=0{
		return 0
	}
	tail:=make([]int,0)
	for _,num:=range nums{
		i:=sort.SearchInts(tail,num)
		if i==len(tail){
			tail=append(tail,num)
		}else{
			tail[i]=num
		}
	}
	return len(tail)
}
// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

