/*
 * @lc app=leetcode.cn id=460 lang=golang
 * @lcpr version=30204
 *
 * [460] LFU 缓存
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type entry struct{
    key ,value ,freq int
}
type LFUCache struct {
    capacity int
    minfreq int 
    keyToNode map[int]*list.Element()
    freqToList map[int]*list.List()
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity:capacity,
        keyToNode: map[int]*list.Element(),
        freqToList: map[int]*list.List(),
    
    }
}
func (this *LFUCache)getEntry(key int)*entry{
    node:=this.keyToNode[key]
    if node==nil{
        return nil
    }
    e:=node.Value.(*entry)
    lst:=this.freqToList[e.freq]
    lst.Remove(node)
    if lst.Len()==0{
        delete(this.freqToList,e.freq)
        if this.minfreq==e.freq{
            this.minfreq++
        }
    }
    e.freq++
    this.PushFront(e)
    return e

}

func (this *LFUCache) Get(key int) int {
    if e:=this.getEntry(key);e!=nil{
        return e.value
    }
    return -1
}

func (this *LFUCache) Put(key int, value int) {
    if e:=this.getEntry(key);e!=nil{
        e.value=value
        return
    
    }
    if len(this.keyToNode)==this.capacity{
        lst:=this.freqToList(this.minfreq)
        delete(this.keyToNode,lst.Remove(lst.Back()).(*entry).key)
        if lst.len()==0{
            delete(this.freqToList,this.minfreq)
        }
        
    }
            
this.PushFront(entry{key,value,1})
this.minfreq=1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"][[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]\n
// @lcpr case=end

*/

