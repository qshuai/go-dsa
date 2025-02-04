package types

import (
	"fmt"
	"math"
	"strings"
)

const bitSize = 64
const maxLeftShift = bitSize - 1

type BitMap struct {
	base uint64
	bits int

	data []uint64
}

// Add adds e element to the bitmap
func (b *BitMap) Add(ele uint64) {
	b.checkBound(ele)
	idx := (ele - b.base) / bitSize
	offset := (ele - b.base) % bitSize
	b.data[idx] |= (1 << (maxLeftShift - offset))
}

// Contain returns true if the bitmap contain the element
func (b *BitMap) Contain(ele uint64) bool {
	b.checkBound(ele)
	idx := (ele - b.base) / bitSize
	offset := (ele - b.base) % bitSize
	target := b.data[idx]
	return target&(1<<(maxLeftShift-offset)) != 0
}

// Remove removes the element from the bitmap
func (b *BitMap) Remove(ele uint64) {
	b.checkBound(ele)
	idx := (ele - b.base) / bitSize
	offset := (ele - b.base) % bitSize
	b.data[idx] &= (math.MaxUint64 ^ (1 << (maxLeftShift - offset)))
}

// String return the binary bit string of the bitmap
func (b *BitMap) String() string {
	if b == nil || len(b.data) <= 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.Grow(b.bits)
	for offset, item := range b.data {
		sub := (offset+1)*bitSize - b.bits
		if sub > 0 {
			sb.WriteString(fmt.Sprintf("%064b", item)[:(bitSize - sub)])
			break
		}

		sb.WriteString(fmt.Sprintf("%064b", item))
	}

	return sb.String()
}

func (b *BitMap) checkBound(ele uint64) {
	if ele > b.base+uint64(b.bits)-1 {
		panic("over max bits size of bitmap")
	}
}

// NewBitMap return a bitmap instance
func NewBitMap(base uint64, length int) *BitMap {
	// mod := length % bitSize
	// nums := length / bitSize
	// if mod != 0 {
	// 	nums += 1
	// }

	// optimize: replace the above code
	nums := (length + bitSize - 1) / bitSize

	return &BitMap{
		base: base,
		bits: length,
		data: make([]uint64, nums),
	}
}
