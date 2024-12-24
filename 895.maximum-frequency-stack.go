/*
 * @lc app=leetcode.cn id=895 lang=golang
 * @lcpr version=30204
 *
 * [895] 最大频率栈
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type FreqStack struct {
	stacks [][]int
	cnt    map[int]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}}
}

func (this *FreqStack) Push(val int) {
	c := this.cnt[val]
	if c == len(this.stacks) {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[c] = append(this.stacks[c], val)
	}
	this.cnt[val]++
}

func (this *FreqStack) Pop() int {
	back := len(this.stacks) - 1
	st := this.stacks[back]
	bk := len(st) - 1
	val := st[bk]
	if bk == 0 {
		this.stacks = this.stacks[:back]
	} else {
		this.stacks[back] = st[:bk]
	}
	this.cnt[val]--
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

/*
// @lcpr case=start
// ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]\n
// @lcpr case=end

*/

