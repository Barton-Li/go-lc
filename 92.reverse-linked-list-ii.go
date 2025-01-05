/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30204
 *
 * [92] 反转链表 II
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	for i := 0; i < left-1; i++ {
		first = first.Next
	}
	cur := first.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = first.Next
		first.Next = next

	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/

