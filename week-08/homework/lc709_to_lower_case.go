package homework

func toLowerCase(s string) string {
	n := len(s)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = s[i]
		if s[i] >= 'A' && s[i] <= 'Z' {
			res[i] += 'z' - 'Z'
		}
	}

	return string(res)
}
