package homework

func lengthOfLastWord(s string) int {
	n := len(s)
	ans := 0

	for n > 0 && s[n-1] == ' ' {
		n--
	}

	for n > 0 && s[n-1] != ' ' {
		ans++
		n--
	}

	return ans
}
