package queue

// MyCircularQueue(k): 构造器，设置队列长度为 k 。
// Front: 从队首获取元素。如果队列为空，返回 -1 。
// Rear: 获取队尾元素。如果队列为空，返回 -1 。
// EnQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// DeQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// IsEmpty(): 检查循环队列是否为空。
// IsFull(): 检查循环队列是否已满。

type MyCircularQueue struct {
	list   []int
	write  int
	read   int
	length int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		list:   make([]int, k),
		length: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.list[this.write%this.length] = value
	this.write++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Front() == -1 {
		return false
	}
	this.read++
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	val := this.list[this.read%this.length]
	// this.read++
	return val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	val := this.list[(this.write+this.length-1)%this.length]
	// this.read = (this.read + 1) % this.length
	return val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.write == this.read

}

func (this *MyCircularQueue) IsFull() bool {
	return this.write > this.read && this.write%this.length == this.read
}
