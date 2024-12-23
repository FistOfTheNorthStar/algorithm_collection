package numberutils

// IsNumberAValidPowerOfBase checks if a number is a valid power of a given base
func IsNumberAValidPowerOfBase(number, base int) bool {
	if number == 0 {
		return true
	}

	for number != 1 {
		if number%base != 0 {
			return false
		}
		number = number / base
	}

	return true
}
