package homework

import "unicode"

func reverseOnlyLetters(s string) string {
	ans := make([]rune, 0)
	r := len(s) - 1

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			for !unicode.IsLetter(rune(s[r])) {
				r--
			}

			ans = append(ans, rune(s[r]))
			r--
		} else {
			ans = append(ans, ch)
		}
	}

	return string(ans)
}
