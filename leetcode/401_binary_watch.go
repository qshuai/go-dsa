package leetcode

import (
	"fmt"

	"github.com/qshuai/go-dsa/utils"
)

// https://leetcode.com/problems/binary-watch/

// readBinaryWatch Given an integer turnedOn which represents the number of LEDs that are currently on,
// return all possible times the watch could represent. You may return the answer in any order.
func readBinaryWatch(turnedOn int) []string {
	if turnedOn < 0 || turnedOn >= 9 {
		return []string{}
	}
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	ret := make([]string, 0)
	for hour := 0; hour < 12; hour++ {
		for minute := 0; minute < 60; minute++ {
			if utils.BitCountUint32(uint32(hour))+utils.BitCountUint32(uint32(minute)) == uint32(turnedOn) {
				if minute < 10 {
					ret = append(ret, fmt.Sprintf("%d:0%d", hour, minute))
				} else {
					ret = append(ret, fmt.Sprintf("%d:%d", hour, minute))
				}
			}
		}
	}

	return ret
}
