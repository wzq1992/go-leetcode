package lessons

func validPalindrome(s string) bool {
	return validPalindromeSub(s, 0, len(s)-1, true)
}

func validPalindromeSub(s string, l, r int, canDel bool) bool {
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			if !canDel {
				return false
			}

			return validPalindromeSub(s, l+1, r, false) || validPalindromeSub(s, l, r-1, false)
		}
	}

	return true
}
