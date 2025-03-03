/*
 * @lc app=leetcode.cn id=417 lang=golang
 * @lcpr version=30204
 *
 * [417] 太平洋大西洋水流问题
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// pacificAtlantic 函数用于查找所有可以流向太平洋和大西洋的单元格坐标。
// 参数：
// - heights: 一个二维整数数组，表示每个单元格的高度。
// 返回值：
// - ans: 一个二维整数数组，包含所有可以流向太平洋和大西洋的单元格坐标。

func pacificAtlantic(heights [][]int) (ans [][]int) {
    // 获取矩阵的高度 m 和宽度 n
    m, n := len(heights), len(heights[0])

    // 初始化两个布尔矩阵，分别记录每个单元格是否可以流向太平洋和大西洋
    pacific := make([][]bool, m)
    atlantic := make([][]bool, m)
    for i := range pacific {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    // 定义深度优先搜索（DFS）函数，用于标记可以从某个起点到达的所有单元格
    var dfs func(int, int, [][]bool)
    dfs = func(x, y int, ocean [][]bool) {
        // 如果当前单元格已经访问过，则直接返回
        if ocean[x][y] {
            return
        }
        // 标记当前单元格为已访问
        ocean[x][y] = true

        // 遍历四个方向（上、下、左、右），递归访问相邻且高度不低于当前单元格的单元格
        for _, d := range dirs {
            nx, ny := x+d.x, y+d.y
            if 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[x][y] {
                dfs(nx, ny, ocean)
            }
        }
    }

    // 从太平洋边界开始进行 DFS，标记所有可以流向太平洋的单元格
    for i := 0; i < m; i++ {
        dfs(i, 0, pacific)
    }
    for j := 1; j < n; j++ {
        dfs(0, j, pacific)
    }

    // 从大西洋边界开始进行 DFS，标记所有可以流向大西洋的单元格
    for i := 0; i < m; i++ {
        dfs(i, n-1, atlantic)
    }
    for j := 0; j < n-1; j++ {
        dfs(m-1, j, atlantic)
    }

    // 遍历整个矩阵，找到既可以流向太平洋又可以流向大西洋的单元格，并将其坐标加入结果列表
    for i, row := range pacific {
        for j, ok := range row {
            if ok && atlantic[i][j] {
                ans = append(ans, []int{i, j})
            }
        }
    }

    // 返回结果列表
    return
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,1],[1,2]]\n
// @lcpr case=end

*/

