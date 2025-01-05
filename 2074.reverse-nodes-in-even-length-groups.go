/*
 * @lc app=leetcode.cn id=2074 lang=golang
 * @lcpr version=30204
 *
 * [2074] 反转偶数长度组的节点
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
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var nodes []*ListNode
	for node, size := head, 1; node != nil; node = node.Next {
		nodes = append(nodes, node)
		if len(nodes) == size || node.Next == nil {
			if n := len(nodes); n%2 == 0 {
				for i := 0; i < n/2; i++ {
					nodes[i].Val, nodes[n-i-1].Val = nodes[n-i-1].Val, nodes[i].Val
				}
			}
			nodes = nil
			size++
		}

	}
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [5,2,6,3,9,1,7,3,8,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,0,6]\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n
// @lcpr case=end

*/

