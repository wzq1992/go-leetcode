package homework

import "strings"

func reverseWords(s string) string {
	l, r := 0, len(s)-1

	// 去除前缀空格
	for l <= r && s[l] == ' ' {
		l++
	}

	// 去除后缀空格
	for l <= r && s[r] == ' ' {
		r--
	}

	word, words := make([]byte, 0), make([]string, 0)
	for l <= r {
		if s[l] == ' ' && len(word) > 0 {
			words = append([]string{string(word)}, words...)
			word = []byte{}
		} else if s[l] != ' ' {
			word = append(word, s[l])
		}

		l++
	}

	words = append([]string{string(word)}, words...)

	return strings.Join(words, " ")
}
