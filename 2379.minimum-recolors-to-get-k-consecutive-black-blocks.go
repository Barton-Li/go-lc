/*
 * @lc app=leetcode.cn id=2379 lang=golang
 * @lcpr version=30204
 *
 * [2379] 得到 K 个黑块的最少涂色次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	if n < k {
		return -1 // 如果 k 大于 blocks 的长度，无法形成连续的 k 个黑色块
	}

	// 初始窗口的操作次数
	operations := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			operations++
		}
	}

	minOps := operations

	// 滑动窗口
	for i := k; i < n; i++ {
		// 移出窗口的元素
		if blocks[i-k] == 'W' {
			operations--
		}
		// 移入窗口的元素
		if blocks[i] == 'W' {
			operations++
		}
		// 更新最小操作次数
		minOps = min(minOps, operations)
	}

	return minOps
}

// @lc code=end

/*
// @lcpr case=start
// "WBBWWBBWBW"\n7\n
// @lcpr case=end

// @lcpr case=start
// "WBWBBBW"\n2\n
// @lcpr case=end

*/

