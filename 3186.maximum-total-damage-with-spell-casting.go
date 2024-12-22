/*
 * @lc app=leetcode.cn id=3186 lang=golang
 * @lcpr version=30204
 *
 * [3186] 施咒的最大总伤害
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumTotalDamage(power []int) int64 {
	count := map[int]int{}
	for _, p := range power {
		count[p]++
	}
	arr := []int{}
	for k := range count {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	n := len(arr)
	f := make([]int64, n+1)
	j := 0
	for i, v := range arr {
		f[i+1] = f[i]
		for arr[j] < v-2 {
			j++
		}
		f[i+1] = max(f[i+1], f[j]+int64(arr[i]*count[arr[i]]))
	}
	return f[n]

}

// @lc code=end

/*
// @lcpr case=start
// [1,1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,1,6,6]\n
// @lcpr case=end

*/

