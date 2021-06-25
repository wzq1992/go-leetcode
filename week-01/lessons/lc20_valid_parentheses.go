package lessons

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	mapStr := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	var stack []byte

	for i := 0; i < length; i++ {
		if mapStr[s[i]] > 0 {
			if len(stack) == 0 || mapStr[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
