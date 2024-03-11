package alg

// This function calculates the greatest common divisor (GCD) of two numbers.
func MaxDiv(first, second int) int {
	max := 0

	for i := 1; i <= first || i <= second; i++ {
		if first%i == 0 && second%i == 0 && i > max {
			max = i
		}
	}

	return max
}
