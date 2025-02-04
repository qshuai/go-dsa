package utils

func Abs[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](num T) T {
	if num < 0 {
		return -num
	}
	return num
}
