package main

func subtractProductAndSum(n int) int {
	s, p := 0, 1
	for n != 0 {
		d := n % 10
		n /= 10

		s += d
		p *= d
	}

	return p - s
}