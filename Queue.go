package main

type Queue struct {
	items []int
}

// 入队
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// 出队
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// 查看对头
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	return item, true
}

// 队列长度
func (q *Queue) Size() int {
	return len(q.items)
}

// 是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
