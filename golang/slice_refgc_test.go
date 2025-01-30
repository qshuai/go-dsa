package golang

import (
	"testing"
	"unsafe"

	"github.com/qshuai/go-dsa/types"
	"github.com/qshuai/go-dsa/utils"
)

func TestSliceRefGc(t *testing.T) {
	capacity := 10
	s := make([]any, capacity)
	sh := (*types.SliceHeader)(unsafe.Pointer(&s))

	ss := s[:0]
	ssh := (*types.SliceHeader)(unsafe.Pointer(&ss))

	// reference the same memory
	utils.AssertTrue(t, sh.Data == ssh.Data)
	utils.AssertEqual(t, cap(ss), capacity)

	ss = append(ss, 1)
	ssh = (*types.SliceHeader)(unsafe.Pointer(&ss))
	utils.AssertTrue(t, sh.Data == ssh.Data)
}
