package internal

import "math"

func MmsModel(lambda, mu, s float64) (L, Lq, W, Wq float64) {
	rho := lambda / (mu * s)
	p0 := 1.0 / (1.0 + sumP0(lambda, mu, s))

	Lq = (math.Pow(rho, s) * p0 * (1.0 - rho)) / (factorial(s-1) * math.Pow(1.0-rho, 2))
	L = Lq + lambda/mu

	Wq = Lq / lambda
	W = Wq + 1.0/mu

	return L, Lq, W, Wq
}

func sumP0(lambda, mu, s float64) float64 {
	sum := 0.0
	rho := lambda / (mu * s)
	for n := 0.0; n <= s-1; n++ {
		sum += math.Pow(rho, n) / factorial(n)
	}
	return sum
}

func factorial(n float64) float64 {
	if n <= 1 {
		return 1.0
	}
	return n * factorial(n-1)
}
