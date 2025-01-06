package leetcode

import "github.com/qshuai/go-dsa/types"

// https://leetcode.com/problems/implement-queue-using-stacks
//
// Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should
// support all the functions of a normal queue (push, peek, pop, and empty).
//
// Implement the MyQueue class:
// void push(int x) Pushes element x to the back of the queue.
// int pop() Removes the element from the front of the queue and returns it.
// int peek() Returns the element at the front of the queue.
// boolean empty() Returns true if the queue is empty, false otherwise.
// Notes:
//
// You must use only standard operations of a stack, which means only push to top, peek/pop from top,
// size, and is empty operations are valid.
// Depending on your language, the stack may not be supported natively. You may simulate a stack using
// a list or deque (double-ended queue) as long as you use only a stack's standard operations.
//
// Examples:
// Input:
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// Output:
// [null, null, null, 1, 1, false]
//
// Explanation:
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
//
// Constraints:
// 1 <= x <= 9
// At most 100 calls will be made to push, pop, peek, and empty.
// All the calls to pop and peek are valid.
//
// Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity?
// In other words, performing n operations will take overall O(n) time even if one of those operations
// may take longer.

type MyQueue[T any] struct {
	stack1 *types.Stack[T]
	stack2 *types.Stack[T]
}

func (mq *MyQueue[T]) Push(x T) {
	mq.stack1.Push(x)
}

func (mq *MyQueue[T]) Pop() T {
	if mq.stack2.Size() > 0 {
		v, _ := mq.stack2.Pop()
		return v
	}

	for mq.stack1.Size() > 0 {
		v, _ := mq.stack1.Pop()
		mq.stack2.Push(v)
	}

	v, _ := mq.stack2.Pop()
	return v
}

func (mq *MyQueue[T]) Peek() T {
	if mq.stack2.Size() > 0 {
		v, _ := mq.stack2.Peek()
		return v
	}

	for mq.stack1.Size() > 0 {
		v, _ := mq.stack1.Pop()
		mq.stack2.Push(v)
	}

	v, _ := mq.stack2.Peek()
	return v
}

func (mq *MyQueue[T]) Empty() bool {
	return mq.stack1.Empty() && mq.stack2.Empty()
}

func NewMyQueue[T any]() *MyQueue[T] {
	return &MyQueue[T]{
		stack1: types.NewStack[T](),
		stack2: types.NewStack[T](),
	}
}
