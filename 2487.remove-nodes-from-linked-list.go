/*
 * @lc app=leetcode.cn id=2487 lang=golang
 * @lcpr version=30204
 *
 * [2487] 从链表中移除节点
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
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	node := removeNodes(head.Next)
	if node.Val > head.Val {
		return node
	}
	head.Next = node
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [5,2,13,3,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n
// @lcpr case=end

*/

