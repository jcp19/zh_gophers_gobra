package heap

func SwapInPlace(a *int, b *int) {
	*a = *a + *b // *a = A + B
	*b = *a - *b // *b = A + B - B = A
	*a = *a - *b // *a = A + B - A = B
}

func client() {
	x, y := new(int), new(int)
	*x, *y = 1, 2
	SwapInPlace(x, y)
}
