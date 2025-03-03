/*
 * @lc app=leetcode.cn id=2181 lang=golang
 * @lcpr version=30204
 *
 * [2181] 合并零之间的节点
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
 */0
func mergeNodes(head *ListNode) *ListNode {
	tail := head
	for curr := head.Next; curr.Next != nil; curr = curr.Next {
		if curr.Val != 0 {
			tail.Val += curr.Val
		} else {
			tail = tail.Next
			tail.Val = 0
		}
	}
	tail.Next = nil
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [0,3,1,0,4,5,2,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,0,2,2,0]\n
// @lcpr case=end

*/

