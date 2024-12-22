/*
 * @lc app=leetcode.cn id=904 lang=golang
 * @lcpr version=30204
 *
 * [904] 水果成篮
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func totalFruit(fruits []int) int {
	n := len(fruits)
	if n <= 2 {
		return n
	}
	fruitCount := make(map[int]int)
	start, maxFruits := 0, 0
	for end := 0; end < n; end++ {
		fruitCount[fruits[end]]++
		for len(fruitCount) > 2 {
			fruitCount[fruits[start]]--
			if fruitCount[fruits[start]] == 0 {
				delete(fruitCount, fruits[start])
			}
			start++
		}
		windowSize := end - start + 1
		if windowSize > maxFruits {
			maxFruits = windowSize
		}
	}
	return maxFruits

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,1,2,1,1,2,3,3,4]\n
// @lcpr case=end

*/

