package lessons

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	n := len(s)
	s = " " + s

	ansLen, ansStart := 0, 0

	// 中心是一个字符
	for center := 1; center <= n; center++ {
		l, r := center-1, center+1
		for l > 0 && r <= n && s[l] == s[r] {
			l--
			r++
		}

		// l + 1 ~ r - 1
		if r-l-1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}

	// 中心是两个字符(或者空)
	for center := 1; center < n; center++ {
		l, r := center, center+1
		for l > 0 && r <= n && s[l] == s[r] {
			l--
			r++
		}

		// l + 1 ~ r - 1
		if r-l-1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}

	return string(s[ansStart : ansStart+ansLen])
}
