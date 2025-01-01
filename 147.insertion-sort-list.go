/*
 * @lc app=leetcode.cn id=147 lang=golang
 * @lcpr version=30204
 *
 * [147] 对链表进行插入排序
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
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val < cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

*/

