package leetcode

/**
 * ************************** circular queue implementation **************************
 */

type MyCircularDeque struct {
	head, tail int
	capacity   int
	arr        []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head:     -1,
		tail:     -1,
		capacity: k,
		arr:      make([]int, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.arr[0] = value
		this.head = 0
		this.tail = 0
		return true
	}

	this.head = (this.head - 1 + this.capacity) % this.capacity
	this.arr[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.arr[0] = value
		this.head = 0
		this.tail = 0
		return true
	}

	this.tail = (this.tail + 1) % this.capacity
	this.arr[this.tail] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	if this.head == (this.tail+1)%this.capacity {
		this.head = -1
		this.tail = -1
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.tail = (this.tail - 1 + this.capacity) % this.capacity
	if this.head == (this.tail+1)%this.capacity {
		this.head = -1
		this.tail = -1
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == -1 && this.tail == -1
}

func (this *MyCircularDeque) IsFull() bool {
	return this.head == (this.tail+1)%this.capacity
}

/**
 * ************************** slice implementation **************************
 */
// type MyCircularDeque struct {
// 	arr      []int
// 	capacity int
// }
//
// func Constructor(k int) MyCircularDeque {
// 	return MyCircularDeque{
// 		arr:      nil,
// 		capacity: k,
// 	}
// }
//
// func (this *MyCircularDeque) InsertFront(value int) bool {
// 	if len(this.arr) >= this.capacity {
// 		return false
// 	}
//
// 	this.arr = append([]int{value}, this.arr...)
// 	return true
// }
//
// func (this *MyCircularDeque) InsertLast(value int) bool {
// 	if len(this.arr) >= this.capacity {
// 		return false
// 	}
//
// 	this.arr = append(this.arr, value)
// 	return true
// }
//
// func (this *MyCircularDeque) DeleteFront() bool {
// 	if len(this.arr) == 0 {
// 		return false
// 	}
//
// 	this.arr = this.arr[1:]
// 	return true
// }
//
// func (this *MyCircularDeque) DeleteLast() bool {
// 	if len(this.arr) == 0 {
// 		return false
// 	}
//
// 	this.arr = this.arr[:len(this.arr)-1]
// 	return true
// }
//
// func (this *MyCircularDeque) GetFront() int {
// 	if len(this.arr) == 0 {
// 		return -1
// 	}
//
// 	return this.arr[0]
// }
//
// func (this *MyCircularDeque) GetRear() int {
// 	if len(this.arr) == 0 {
// 		return -1
// 	}
//
// 	return this.arr[len(this.arr)-1]
// }
//
// func (this *MyCircularDeque) IsEmpty() bool {
// 	return len(this.arr) == 0
// }
//
// func (this *MyCircularDeque) IsFull() bool {
// 	return len(this.arr) == this.capacity
// }

/**
 * ************************** Doubly linked list implementation **************************
 */
//
// type MyDoublyListNode struct {
// 	value int
// 	prev  *MyDoublyListNode
// 	next  *MyDoublyListNode
// }
//
// type MyCircularDeque struct {
// 	size, capacity int
// 	head           *MyDoublyListNode
// 	tail           *MyDoublyListNode
// }
//
// func Constructor(k int) MyCircularDeque {
// 	return MyCircularDeque{capacity: k}
// }
//
// func (this *MyCircularDeque) InsertFront(value int) bool {
// 	if this.size >= this.capacity {
// 		return false
// 	}
//
// 	this.head = &MyDoublyListNode{
// 		value: value,
// 		prev:  nil,
// 		next:  this.head,
// 	}
// 	if this.head.next != nil {
// 		this.head.next.prev = this.head
// 	}
// 	if this.size == 0 {
// 		this.tail = this.head
// 	}
// 	this.size++
//
// 	return true
// }
//
// func (this *MyCircularDeque) InsertLast(value int) bool {
// 	if this.size >= this.capacity {
// 		return false
// 	}
//
// 	this.tail = &MyDoublyListNode{
// 		value: value,
// 		prev:  this.tail,
// 		next:  nil,
// 	}
// 	if this.tail.prev != nil {
// 		this.tail.prev.next = this.tail
// 	}
// 	if this.size == 0 {
// 		this.head = this.tail
// 	}
// 	this.size++
//
// 	return true
// }
//
// func (this *MyCircularDeque) DeleteFront() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
//
// 	this.head = this.head.next
// 	this.size--
// 	return true
// }
//
// func (this *MyCircularDeque) DeleteLast() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
//
// 	this.tail = this.tail.prev
// 	this.size--
// 	return true
// }
//
// func (this *MyCircularDeque) GetFront() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
//
// 	return this.head.value
// }
//
// func (this *MyCircularDeque) GetRear() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
//
// 	return this.tail.value
// }
//
// func (this *MyCircularDeque) IsEmpty() bool {
// 	return this.size == 0
// }
//
// func (this *MyCircularDeque) IsFull() bool {
// 	return this.size == this.capacity
// }
