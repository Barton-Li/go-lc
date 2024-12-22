/*
 * @lc app=leetcode.cn id=1847 lang=golang
 * @lcpr version=30204
 *
 * [1847] 最近的房间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func closestRoom(rooms [][]int, queries [][]int) []int {
	n := len(rooms)
	k := len(queries)
	answer := make([]int, k)

	for j, query := range queries {
		preferred := query[0]
		minSize := query[1]
		minDiff := -1
		roomId := -1

		for i := 0; i < n; i++ {
			id := rooms[i][0]
			size := rooms[i][1]
			if size >= minSize {
				diff := abs(id - preferred)
				if minDiff == -1 || diff < minDiff || (diff == minDiff && id < roomId) {
					minDiff = diff
					roomId = id
				}
			}
		}

		answer[j] = roomId
	}

	return answer
}
func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[2,2],[1,2],[3,2]]\n[[3,1],[3,3],[5,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[2,3],[3,5],[4,1],[5,2]]\n[[2,3],[2,4],[2,5]]\n
// @lcpr case=end

*/

