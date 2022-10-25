package intro

func Abs(x int32) int32 {
	if x >= 0 {
		return x
	}
	return -1 * x
}

















































































// ##(--overflow)
// @ ghost const MIN_INT32 = -2147483648