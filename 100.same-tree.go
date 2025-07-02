/*
 * @lc app=leetcode.cn id=100 lang=golang
 * @lcpr version=30204
 *
 * [100] 相同的树
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil&&q=nil{
		return true
	}
	if p==nil||q==nil{
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return isSameTree(q.Left,p.Left)&&isSameTree(q.Right，p.Right)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1]\n[1,1,2]\n
// @lcpr case=end

*/

