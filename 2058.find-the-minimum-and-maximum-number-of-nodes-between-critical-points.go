/*
 * @lc app=leetcode.cn id=2058 lang=golang
 * @lcpr version=30204
 *
 * [2058] 找出临界点之间的最小和最大距离
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
func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next
	first, last, minDis := 0, 0, math.MaxInt32
	for i, prev := 1, 0; c != nil; i++ {
		if a.Val < b.Val && b.Val > c.Val || a.Val > b.Val && b.Val < c.Val {
			if first == 0 {
				first = i
			}
			last = i
			if prev > 0 && i-prev < minDis {
				minDis = i - prev
			}
			prev = i
		}
		a, b, c = b, c, c.Next
	}
	if first == last {
		return []int{-1, -1}
	}
	return []int{minDis, last - first}
}

// @lc code=end

/*
// @lcpr case=start
// [3,1]\n
// @lcpr case=end

// @lcpr case=start
// [5,3,1,2,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,2,3,2,2,2,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,3,2]\n
// @lcpr case=end

*/

