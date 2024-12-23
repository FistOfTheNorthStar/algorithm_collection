package algorithms

// IsPrimeNumber checks if the given number is prime
func IsPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	}

	for div := (number / 2) + 1; div > 1; div-- {
		if number%div == 0 {
			return false
		}
	}

	return true
}
