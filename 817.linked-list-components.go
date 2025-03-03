/*
 * @lc app=leetcode.cn id=817 lang=golang
 * @lcpr version=30204
 *
 * [817] 链表组件
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
func numComponents(head *ListNode, nums []int) int {
	// 将 nums 中的元素存入集合 set，用于快速查找
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	// 记录连通分量的数量
	n := 0

	// 遍历链表并计算连通分量的数量
	for inSet := false; head != nil; head = head.Next {
		// 如果当前节点值不在 set 中，重置 inSet 标志
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			// 如果当前节点值在 set 中且之前不在 set 中，则发现一个新的连通分量
			inSet = true
			n++
		}
	}

	return n
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,2,3]\n[0,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,3,4]\n[0,3,1,4]\n
// @lcpr case=end

*/

