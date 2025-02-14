package leetcode

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/meeting-rooms-iii
//
// You are given an integer n. There are n rooms numbered from 0 to n - 1.
// You are given a 2D integer array meetings where meetings[i] = [starti, endi] means that a meeting will be held during the half-closed time interval [starti, endi). All the values of starti are unique.
// Meetings are allocated to rooms in the following manner:
// Each meeting will take place in the unused room with the lowest number.
// If there are no available rooms, the meeting will be delayed until a room becomes free. The delayed meeting should have the same duration as the original meeting.
// When a room becomes unused, meetings that have an earlier original start time should be given the room.
// Return the number of the room that held the most meetings. If there are multiple rooms, return the room with the lowest number.
// A half-closed interval [a, b) is the interval between a and b including a and not including b.
//
// Example 1:
// Input: n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]
// Output: 0
// Explanation:
// - At time 0, both rooms are not being used. The first meeting starts in room 0.
// - At time 1, only room 1 is not being used. The second meeting starts in room 1.
// - At time 2, both rooms are being used. The third meeting is delayed.
// - At time 3, both rooms are being used. The fourth meeting is delayed.
// - At time 5, the meeting in room 1 finishes. The third meeting starts in room 1 for the time period [5,10).
// - At time 10, the meetings in both rooms finish. The fourth meeting starts in room 0 for the time period [10,11).
// Both rooms 0 and 1 held 2 meetings, so we return 0.
//
// Example 2:
// Input: n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
// Output: 1
// Explanation:
// - At time 1, all three rooms are not being used. The first meeting starts in room 0.
// - At time 2, rooms 1 and 2 are not being used. The second meeting starts in room 1.
// - At time 3, only room 2 is not being used. The third meeting starts in room 2.
// - At time 4, all three rooms are being used. The fourth meeting is delayed.
// - At time 5, the meeting in room 2 finishes. The fourth meeting starts in room 2 for the time period [5,10).
// - At time 6, all three rooms are being used. The fifth meeting is delayed.
// - At time 10, the meetings in rooms 1 and 2 finish. The fifth meeting starts in room 1 for the time period [10,12).
// Room 0 held 1 meeting while rooms 1 and 2 each held 2 meetings, so we return 1.
//
// Constraints:
// 1 <= n <= 100
// 1 <= meetings.length <= 10^5
// meetings[i].length == 2
// 0 <= starti < endi <= 5 * 10^5
// All the values of starti are unique.

type freeRoom []int

func (f *freeRoom) Len() int {
	return len(*f)
}

func (f *freeRoom) Less(i, j int) bool {
	return (*f)[i] < (*f)[j]
}

func (f *freeRoom) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *freeRoom) Push(x any) {
	*f = append(*f, x.(int))
}

func (f *freeRoom) Pop() any {
	last := (*f)[len(*f)-1]
	*f = (*f)[:len(*f)-1]
	return last
}

type busyRoom [][]int

func (b *busyRoom) Len() int {
	return len(*b)
}

func (b *busyRoom) Less(i, j int) bool {
	if (*b)[i][0] == (*b)[j][0] {
		return (*b)[i][1] < (*b)[j][1]
	}

	return (*b)[i][0] < (*b)[j][0]
}

func (b *busyRoom) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *busyRoom) Push(x any) {
	*b = append(*b, x.([]int))
}

func (b *busyRoom) Pop() any {
	last := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return last
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// 小顶堆
	free := &freeRoom{}
	for i := 0; i < n; i++ {
		heap.Push(free, i)
	}
	busy := &busyRoom{}
	booked := make([]int, n)

	for i := 0; i < len(meetings); i++ {
		for busy.Len() > 0 && (*busy)[0][0] <= meetings[i][0] {
			idx := (*busy)[0][1]
			heap.Pop(busy)
			heap.Push(free, idx)
		}

		if free.Len() > 0 {
			idx := (*free)[0]
			heap.Pop(free)
			heap.Push(busy, []int{meetings[i][1], idx})
			booked[idx]++
		} else {
			removed := heap.Pop(busy).([]int)
			heap.Push(busy, []int{removed[0] + meetings[i][1] - meetings[i][0], removed[1]})
			booked[removed[1]]++
		}
	}

	var maxIdx, maxBooked int
	for idx, nums := range booked {
		if nums > maxBooked {
			maxBooked = nums
			maxIdx = idx
		}
	}

	return maxIdx
}
