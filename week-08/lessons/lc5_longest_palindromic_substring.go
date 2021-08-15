package lessons

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	n := len(s)
	s = " " + s
	p := int64(1e9) + 7

	// 预处理 hash
	preH, sufH, p131 := make([]int64, n+1), make([]int64, n+2), make([]int64, n+1)
	p131[0] = 1
	for i := 1; i <= n; i++ {
		j := n - i + 1
		preH[i] = (preH[i-1]*131 + (int64)(s[i]-'a'+1)) % p
		sufH[j] = (sufH[j+1]*131 + (int64)(s[j]-'a'+1)) % p
		p131[i] = p131[i-1] * 131 % p
	}

	var calcPre func(l, r int) int64
	calcPre = func(l, r int) int64 {
		return ((preH[r]-preH[l-1]*p131[r-l+1])%p + p) % p
	}

	var calcSuf func(l, r int) int64
	calcSuf = func(l, r int) int64 {
		return ((sufH[l]-sufH[r+1]*p131[r-l+1])%p + p) % p
	}

	var isPalindrome func(l, r int) bool
	isPalindrome = func(l, r int) bool {
		if l < 1 || r > n {
			return false
		}
		if l > r {
			return true
		}

		return calcPre(l, r) == calcSuf(l, r)
	}

	ansLen, ansStart := 0, 0

	// 中心是一个字符
	for center := 1; center <= n; center++ {
		lenL, lenR := 0, n
		for lenL < lenR {
			length := (lenL + lenR + 1) / 2
			l, r := center-length, center+length
			if isPalindrome(l, r) {
				lenL = length
			} else {
				lenR = length - 1
			}
		}

		// 两侧最多扩展 lenL 再加一个中心
		if lenL*2+1 > ansLen {
			ansLen = lenL*2 + 1
			ansStart = center - lenL
		}
	}

	// 中心是两个字符(或者空)
	for center := 1; center < n; center++ {
		lenL, lenR := 0, n
		for lenL < lenR {
			length := (lenL + lenR + 1) / 2
			l, r := center-length+1, center+length
			if isPalindrome(l, r) {
				lenL = length
			} else {
				lenR = length - 1
			}
		}

		// 两侧分别是 l ~ center, center + 1 ~ r
		if lenL*2 > ansLen {
			ansLen = lenL * 2
			ansStart = center - lenL + 1
		}
	}

	return string(s[ansStart : ansStart+ansLen])
}
