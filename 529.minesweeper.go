/*
 * @lc app=leetcode.cn id=529 lang=golang
 * @lcpr version=30204
 *
 * [529] 扫雷游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// dirX 和 dirY 用于定义八个可能的搜索方向
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

// updateBoard 是解决扫雷问题的主要函数
// board 代表雷区的当前状态，click 是用户点击的位置
// 返回更新后的雷区状态
func updateBoard(board [][]byte, click []int) [][]byte {
    x, y := click[0], click[1]
    // 如果点击的位置是雷，直接将其标记为 X
    if board[x][y] == 'M' {
        board[x][y] = 'X'
    } else {
        // 否则，从点击的位置开始进行深度优先搜索
        dfs(board, x, y)
    }
    return board
}

// dfs 是用于更新雷区的深度优先搜索函数
// board 代表雷区的当前状态，x 和 y 是当前搜索的位置
func dfs(board [][]byte, x, y int) {
    // 初始化雷的计数器
    cnt := 0
    // 遍历八个方向，统计周围的雷数量
    for i := 0; i < 8; i++ {
        tx, ty := x+dirX[i], y+dirY[i]
        // 忽略越界的坐标
        if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
            continue
        }
        // 如果周围有雷，增加计数
        if board[tx][ty] == 'M' {
            cnt++
        }
    }
    // 如果周围有雷，将当前位置更新为雷的数量
    if cnt > 0 {
        board[x][y] = byte(cnt + '0')
    } else {
        // 否则，将当前位置标记为已搜索，并继续搜索周围的位置
        board[x][y] = 'B'
        for i := 0; i < 8; i++ {
            tx, ty := x+dirX[i], y+dirY[i]
            // 忽略越界或已经搜索过的位置
            if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
                continue
            }
            // 继续深度优先搜索
            dfs(board, tx, ty)
        }
    }
}

// @lc code=end

/*
// @lcpr case=start
// [["E","E","E","E","E"],["E","E","M","E","E"],["E","E","E","E","E"],["E","E","E","E","E"]]\n[3,0]\n
// @lcpr case=end

// @lcpr case=start
// [["B","1","E","1","B"],["B","1","M","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]]\n[1,2]\n
// @lcpr case=end

*/

