/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=30204
 *
 * [146] LRU 缓存
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type entry struct {
	key, value int
}
type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

// Get 从缓存中获取指定的元素。如果元素不存在，则返回-1。
// 如果元素存在，将该元素移动到列表的前面，并返回其值。
func (this *LRUCache) Get(key int) int {
	// 查找键对应的节点
	node := this.keyToNode[key]
	// 如果节点不存在，返回-1
	if node == nil {
		return -1
	}
	// 将访问的节点移动到列表的前面
	this.list.MoveToFront(node)
	// 返回节点的值
	return node.Value.(entry).value
}

// Put 将指定的键值对添加到缓存中。
// 如果键已经存在，则更新其值并将该节点移动到列表的前面。
// 如果键不存在，则在列表的前面添加一个新的节点。
// 如果缓存超过了其容量，则删除列表最后的节点。
func (this *LRUCache) Put(key int, value int) {
	// 如果键已经存在，更新其值并将其移动到列表的前面
	if node := this.keyToNode[key]; node != nil {
		node.Value = entry{key, value}
		this.list.MoveToFront(node)
		return
	}
	// 如果键不存在，添加新的节点到列表的前面
	this.keyToNode[key] = this.list.PushFront(entry{key, value})
	// 如果缓存超过了其容量，删除列表最后的节点
	if len(this.keyToNode) > this.capacity {
		delete(this.keyToNode, this.list.Remove(this.list.Back()).(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end



