/*
 * @lc app=leetcode.cn id=1669 lang=golang
 * @lcpr version=30204
 *
 * [1669] 合并两个链表
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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prevA := list1
	for i := 0; i < a-1; i++ {
		prevA = prevA.Next
	}
	prevB := prevA
	for i := 0; i < b-a+2; i++ {
		prevB = prevB.Next
	}
	prevA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = prevB
	return list1
}

// @lc code=end

/*
// @lcpr case=start
// [10,1,13,6,9,5]\n3\n4\n[1000000,1000001,1000002]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,3,4,5,6]\n2\n5\n[1000000,1000001,1000002,1000003,1000004]\n
// @lcpr case=end

*/

