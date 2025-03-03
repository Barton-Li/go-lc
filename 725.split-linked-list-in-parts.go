/*
 * @lc app=leetcode.cn id=725 lang=golang
 * @lcpr version=30204
 *
 * [725] 分隔链表
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
func splitListToParts(head *ListNode, k int) []*ListNode {
	// 计算链表的总长度
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	// 计算每个部分的基础长度和需要多分配一个节点的部分数量
	q, r := n/k, n%k

	// 创建一个长度为 k 的切片用于存储结果
	parts := make([]*ListNode, k)

	// 遍历链表并分割成 k 个部分
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr // 当前部分的起始节点

		// 计算当前部分的长度
		partSize := q
		if i < r {
			partSize++ // 前 r 个部分需要多分配一个节点
		}

		// 移动 curr 到当前部分的末尾
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}

		// 断开当前部分与下一个部分的连接
		curr, curr.Next = curr.Next, nil
	}

	return parts
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n3\n
// @lcpr case=end

*/

