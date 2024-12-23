package numberutils

import (
	"fmt"
	"strings"
)

const (
	BUZZ = "Buzz"
	FIZZ = "Fizz"
)

// FizzBuzz returns a string with numbers from 1 to N, replacing multiples of 3 with Fizz,
// multiples of 5 with Buzz, and multiples of both with Fizz-Buzz
func FizzBuzz(N int) string {
	var builder strings.Builder

	for i := 1; i <= N; i++ {
		switch {
		case i%15 == 0:
			builder.WriteString(FIZZ + "-" + BUZZ)
		case i%3 == 0:
			builder.WriteString(FIZZ)
		case i%5 == 0:
			builder.WriteString(BUZZ)
		default:
			builder.WriteString(fmt.Sprintf("%d", i))
		}

		if i < N {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func PrintFizzBuzz(N int) {
	fmt.Println(FizzBuzz(N))
}
