package heap

func SwapInPlace(a *int, b *int) {
	*a = *a + *b // *a = A + B
	*b = *a - *b // *b = A + B - B = A
	*a = *a - *b // *a = A + B - A = B
}









