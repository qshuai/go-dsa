package leetcode

import (
	"golang.org/x/exp/constraints"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/min-stack/
//
// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// Implement the MinStack class:
// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.
//
// Example 1:
//
// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// Output
// [null,null,null,null,-3,null,0,-2]
//
// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2
//
// Constraints:
// -2^31 <= val <= 2^31 - 1
// Methods pop, top and getMin operations will always be called on non-empty stacks.
// At most 3 * 10^4 calls will be made to push, pop, top, and getMin.

// 解题思路：栈是一个支持有限制操作(push、pop)的数据容器，由于集合元素添加、删除的约束性，只需要维护对应元素集合的最小值即可。

type MinStack[T constraints.Ordered] struct {
	data *types.Stack[T]
	min  *types.Stack[T]
}

func (ms *MinStack[T]) Push(val T) {
	ms.data.Push(val)

	if ms.min.Empty() {
		ms.min.Push(val)
		return
	}

	minim, _ := ms.min.Peek()
	if val <= minim {
		ms.min.Push(val)
	} else {
		ms.min.Push(minim)
	}
}

func (ms *MinStack[T]) Pop() {
	ms.data.Pop()
	ms.min.Pop()
}

func (ms *MinStack[T]) Top() T {
	v, _ := ms.data.Peek()
	return v
}

func (ms *MinStack[T]) GetMin() T {
	v, _ := ms.min.Peek()
	return v
}

func NewMinStack[T constraints.Ordered]() *MinStack[T] {
	return &MinStack[T]{
		data: types.NewStack[T](),
		min:  types.NewStack[T](),
	}
}
