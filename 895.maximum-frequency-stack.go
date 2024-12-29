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
	stacks [][]int // 用于存储不同频率的元素的栈，stacks[i]存储频率为i+1的元素
	cnt    map[int]int // 用于记录每个元素出现的频率
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}} // 初始化FreqStack，cnt初始化为空map
}

func (this *FreqStack) Push(val int) {
	c := this.cnt[val] // 获取元素val的当前频率
	if c == len(this.stacks) {
		this.stacks = append(this.stacks, []int{val}) // 如果当前频率等于stacks的长度，说明需要增加一个新的频率栈，将val放入新的栈中
	} else {
		this.stacks[c] = append(this.stacks[c], val) // 否则，将val放入对应频率的栈中
	}
	this.cnt[val]++ // 更新元素val的频率
}

func (this *FreqStack) Pop() int {
	back := len(this.stacks) - 1 // 获取频率最高的栈的索引
	st := this.stacks[back] // 获取频率最高的栈
	bk := len(st) - 1 // 获取栈顶元素的索引
	val := st[bk] // 获取栈顶元素
	if bk == 0 {
		this.stacks = this.stacks[:back] // 如果栈中只有一个元素，弹出后删除该栈
	} else {
		this.stacks[back] = st[:bk] // 否则，弹出栈顶元素并更新栈
	}
	this.cnt[val]-- // 更新元素val的频率
	return val // 返回弹出的元素
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

