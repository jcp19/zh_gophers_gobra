package tutorial

// @ requires n >= 0
// @ ensures sum == n * (n+1)/2
// @ decreases
func SumToN(n int) (sum int) {
	sum = 0
	// @ invariant 0 <= i && i <= n+1
	// @ invariant sum == i * (i-1)/2
	// @ decreases n-i
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

// @ requires y > 0
func client(y int) {
	x := SumToN(y)
	_ = x
	// @ assert x == y * (y+1)/2
}
