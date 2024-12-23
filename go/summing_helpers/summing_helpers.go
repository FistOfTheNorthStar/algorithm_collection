package numberutils

// SumOfEvenNumbersWithFor calculates sum of even numbers up to N using a for loop
func SumOfEvenNumbersWithFor(N int) int {
	sum := 0
	for number := 1; number <= N; number++ {
		if number%2 == 0 {
			sum += number
		}
	}
	return sum
}

// SumOfEvenNumbersWithStep calculates sum of even numbers up to N using step-by-2
func SumOfEvenNumbersWithStep(N int) int {
	sum := 0
	for number := 2; number <= N; number += 2 {
		sum += number
	}
	return sum
}
