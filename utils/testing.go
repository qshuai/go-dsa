package utils

import "testing"

// PanicHelper wraps panic branch
func PanicHelper(t *testing.T, fn func(), shouldPanic bool) {
	if shouldPanic {
		ShouldPanic(t, fn)
		return
	}

	ShouldNotPanic(t, fn)
}

// ShouldPanic will report test error if not panic
func ShouldPanic(t *testing.T, fn func()) {
	defer func() {
		_ = recover()
	}()

	fn()

	t.Errorf("should panic")
}

// ShouldNotPanic will report test error if panic
func ShouldNotPanic(t *testing.T, fn func()) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("should not panic")
		}
	}()

	fn()
}
