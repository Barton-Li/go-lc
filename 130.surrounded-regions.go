/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=30204
 *
 * [130] 被围绕的区域
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// solve函数用于解决包围区域问题。它接收一个二维字节切片board作为输入，其中'X'表示围墙，'O'表示待填充的区域。
// 该函数的目标是填充所有被'X'围墙包围的'O'区域，即把这些'O'变成'X'，而不影响与边界相连的'O'区域。
func solve(board [][]byte) {
    // 获取board的行数和列数
    m, n := len(board), len(board[0])

    // dfs是一个深度优先搜索函数，用于标记从(i, j)开始的连通'O'区域，标记为'A'。
    var dfs func(i, j int)
    dfs = func(i, j int) {
        // 如果当前位置(i, j)超出边界或不是'O'，则返回
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
            return
        }
        // 标记当前位置为'A'，表示这个'O'是与边界连通的，不应该被填充
        board[i][j] = 'A'
        // 递归调用dfs，标记所有与当前位置连通的'O'区域
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    // 遍历所有边界，对边界上的'O'进行dfs标记
    for i := 0; i < m; i++ {
        dfs(i, 0)
        dfs(i, n-1)
    }
    for j := 1; j < n-1; j++ {
        dfs(0, j)
        dfs(m-1, j)
    }

    // 遍历整个board，根据标记结果进行填充
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // 如果当前位置被标记为'A'，说明它是与边界连通的'O'，恢复为'O'
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                // 如果当前位置仍然是'O'，说明它被包围，填充为'X'
                board[i][j] = 'X'
            }
        }
    }
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/

