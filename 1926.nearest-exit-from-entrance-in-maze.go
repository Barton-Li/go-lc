/*
 * @lc app=leetcode.cn id=1926 lang=golang
 * @lcpr version=30204
 *
 * [1926] 迷宫中离入口最近的出口
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type pair struct{x,y int}
var dirs=[]pair{{-1,0},{1,0},{0,-1},{0,1}}}
func nearestExit(maze [][]byte, entrance []int) int {
	m,n:=len(maze),len(maze[0])
	s:=pair{entrance[0],entrance[1]}
	maze[s.x][s.y]=0
	q:=[]pair{s}
	for ans:=1;len(q)>0;ans++{
		tmp:=q
		q=nil
		for _,p:=range tmp{
			for _,d:=range dirs{
				if  x,y:=p.x+d.x,p.y+d.y;x>=0&&x<m&&y>=0&&y<n&&maze[x][y]=='.'{
					if x==0||y==0||x==m-1||y==n-1{
						return ans
					}
					maze[x][y]=0
					q=append(q,pair{x,y})
				}
			}
		}
	}
	return -1
}
// @lc code=end



/*
// @lcpr case=start
// [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]]\n[1,2]\n
// @lcpr case=end

// @lcpr case=start
// [["+","+","+"],[".",".","."],["+","+","+"]]\n[1,0]\n
// @lcpr case=end

// @lcpr case=start
// [[".","+"]]\n[0,0]\n
// @lcpr case=end

 */

