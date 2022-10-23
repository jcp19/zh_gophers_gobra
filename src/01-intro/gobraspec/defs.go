package gobraspec

const MIN_INT32 = -2147483648

type OverflowError struct{}

func (o OverflowError) Error() string {
	return "Operation `abs` overflowed"
}
