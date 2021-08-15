package lessons

import "strings"

func isPalindrome(s string) bool {
	t := ""
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= '0' && ch <= '9' ||
			ch >= 'a' && ch <= 'z' {
			t += string(ch)
		}
	}

	l, r := 0, len(t)-1
	for l < r {
		if t[l] != t[r] {
			return false
		}
		l++
		r--
	}

	return true
}
