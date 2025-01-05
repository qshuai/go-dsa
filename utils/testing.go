package utils

import (
	"reflect"
	"testing"
)

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

func AssertEqual(t *testing.T, got, want any) {
	if reflect.DeepEqual(got, want) {
		return
	}

	t.Fatalf("%s want: %v but got: %v", t.Name(), want, got)
}

func AssertTrue(t *testing.T, res bool) {
	if res {
		return
	}

	t.Fatalf("%s want true but got false", t.Name())
}

func AssertFalse(t *testing.T, res bool) {
	if !res {
		return
	}

	t.Fatalf("%s want false but got true", t.Name())
}
