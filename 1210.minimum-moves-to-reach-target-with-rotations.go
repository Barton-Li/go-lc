/*
 * @lc app=leetcode.cn id=1210 lang=golang
 * @lcpr version=30204
 *
 * [1210] 穿过迷宫的最少移动次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type tuple struct{ x, y, s int }

var dirs = []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

// minimumMoves 计算蛇从左上角移动到右下角所需的最小步数。
// g 是一个二维数组，表示网格，其中0表示空地，1表示障碍物。
// 蛇的初始位置是水平占据左上角的两个单元格。
// 如果无法到达右下角，则返回-1。
func minimumMoves(g [][]int) int {
    // n 表示网格的大小。
    n := len(g)
    // vis 用于记录每个位置在水平和垂直状态下是否已被访问，以避免重复计算。
    vis := make([][][2]bool, n)
    for i := range vis {
        vis[i] = make([][2]bool, n)
    }
    vis[0][0][0] = true // 初始位置
    // q 用作队列，存储待处理的位置。
    q := []tuple{{}}
   
    // 使用BFS遍历所有可能的移动方式，直到找到最短路径。
    for step := 1; len(q) > 0; step++ {
        tmp := q
        q = nil
        for _, t := range tmp {
            for _, d := range dirs {
                x, y, s := t.x+d.x, t.y+d.y, t.s^d.s
                x2, y2 := x+s, y+(s^1) // 蛇头
                // 检查新位置是否有效，是否已经访问过，以及是否有障碍物。
                if x2 < n && y2 < n && !vis[x][y][s] &&g[x][y] == 0 && g[x2][y2] == 0 && (d.s == 0 || g[x+1][y+1] == 0) {
                    // 如果蛇的尾巴到达右下角的前一个位置，说明蛇头已经到达右下角。
                    if x == n-1 && y == n-2 {
                        return step
                    }
                    vis[x][y][s] = true
                    // 将新位置添加到队列中，以便后续处理。
                    q = append(q, tuple{x, y, s})
                }
            }
        }
    }
    // 如果无法到达右下角，返回-1。
    return -1
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0,0,0,1],[1,1,0,0,1,0],[0,0,0,0,1,1],[0,0,1,0,1,0],[0,1,1,0,0,0],[0,1,1,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,1,1,1,1],[0,0,0,0,1,1],[1,1,0,0,0,1],[1,1,1,0,0,1],[1,1,1,0,0,1],[1,1,1,0,0,0]]\n
// @lcpr case=end

*/

