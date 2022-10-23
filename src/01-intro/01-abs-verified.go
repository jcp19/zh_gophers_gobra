package intro

// ##(-I .)
// ##(--overflow)

// @ import "gobraspec"

// @ requires x != gobraspec.MIN_INT32
// @ ensures  x >= 0 ==> res == x
// @ ensures  x < 0  ==> res == -x
func Abs(x int32) (res int32) {
	if x >= 0 {
		return x
	}
	return -1 * x
}
