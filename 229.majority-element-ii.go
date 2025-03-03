/*
 * @lc app=leetcode.cn id=229 lang=golang
 * @lcpr version=30204
 *
 * [229] 多数元素 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func majorityElement(nums []int) []int {
// 	count := map[int]int{}
// 	for _, num := range nums {
// 		count[num]++

// 	}
// 	res := []int{}
// 	for num, c := range count {
// 		if c > len(nums)/3 {
// 			res = append(res, num)
// 		}
// 	}
// 	return res
// }
func majorityElement(nums []int) []int {
	res := []int{}
	cand1 := nums[0]
	count1 := 0
	cand2 := nums[0]
	count2 := 0
	for i := 0; i < len(nums); i++ {
		if cand1 == nums[i] {
			cand1++
			continue
		}
		if cand2 == nums[i] {
			cand2++
			continue
		}
		if count1 == 0 {
			cand1 = nums[i]
			count1++
			continue

		}
		if count2 == 0 {
			cand2 = nums[i]
			count2++
			continue
		}
		count1--
		count2--
	}
	count1, count2 = 0, 0
	for i := 0; i < len(nums); i++ {
		if cand1 == nums[i] {
			count1++
		} else if cand2 == nums[i] {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

