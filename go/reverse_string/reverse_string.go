package algorithms

import "errors"

func Reverse(text string) (string, error) {
	if len(text) == 0 {
		return "", errors.New("text is empty")
	}

	chars := []rune(text)
	length := len(chars)

	for i := 0; i < length/2; i++ {
		chars[i], chars[length-1-i] = chars[length-1-i], chars[i]
	}

	return string(chars), nil
}
