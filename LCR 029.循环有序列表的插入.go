/*
 * @lc app=leetcode.cn id=LCR 029 lang=golang
 * @lcpr version=30204
 *
 * [LCR 029] 循环有序列表的插入
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}
	if aNode == nil {
		node.Next = node
		return node
	}
	if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
		return aNode
	}
	cur, next := aNode, aNode.Next
	for next != aNode {
		if x >= cur.Val && x <= next.Val {
			break
		}
		if cur.Val > next.Val {
			if x > cur.Val || x < next.Val {
				break
			}
		}
		cur = cur.Next
		next = next.Next
	}
	cur.Next = node
	node.Next = next
	return aNode
}

// @lc code=end

/*
// @lcpr case=start
// [3,4,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// []\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

