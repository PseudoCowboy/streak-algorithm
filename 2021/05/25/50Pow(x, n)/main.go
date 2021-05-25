package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 0.0
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

// 没想到会超时
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	sum := x
	if n > 0 {
		for i := 0; i < n-1; i++ {
			sum *= x
		}
		return sum
	}
	if n < 0 {
		for i := 0; i < -n-1; i++ {
			sum *= x
		}
		return 1 / sum
	}

	return 0.0
}
