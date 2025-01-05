/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30204
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for curr := head; curr != nil; curr = curr.Next {
		n++
	}
	dummy := &ListNode{Next: head}
	p := dummy
	var prev, curr *ListNode = nil, p.Next
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		next := p.Next
		p.Next.Next = curr
		p.Next = prev
		p = next
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/

