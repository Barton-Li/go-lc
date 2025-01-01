/*
 * @lc app=leetcode.cn id=2807 lang=golang
 * @lcpr version=30204
 *
 * [2807] 在链表中插入最大公约数
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
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for cur := head; cur.Next != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
	}
	return head
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [18,6,10,3]\n
// @lcpr case=end

// @lcpr case=start
// [7]\n
// @lcpr case=end

*/

