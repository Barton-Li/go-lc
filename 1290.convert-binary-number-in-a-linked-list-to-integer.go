/*
 * @lc app=leetcode.cn id=1290 lang=golang
 * @lcpr version=30204
 *
 * [1290] 二进制链表转整数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next

	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,0]\n
// @lcpr case=end

*/

