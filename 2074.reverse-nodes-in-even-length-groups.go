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
    // 定义一个切片用于存储当前组的节点
	var nodes []*ListNode
	// 遍历链表，node 为当前节点，size 为当前组的大小
	for node, size := head, 1; node != nil; node = node.Next {
		// 将当前节点添加到 nodes 列表中
		nodes = append(nodes, node)
		// 如果当前组的节点数等于组大小或已经到达链表末尾
		if len(nodes) == size || node.Next == nil {
			// 获取当前组的节点数量 n
			if n := len(nodes); n%2 == 0 {
				// 如果当前组的节点数量为偶数，则反转该组节点的值
				for i := 0; i < n/2; i++ {
					// 交换对称位置的节点值
					nodes[i].Val, nodes[n-i-1].Val = nodes[n-i-1].Val, nodes[i].Val
				}
			}
			// 清空 nodes 列表，准备处理下一组节点
			nodes = nil
			// 增加组大小
			size++
		}
	}
	// 返回处理后的链表头节点
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

