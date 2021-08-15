package lessons

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}

	haystack = " " + haystack
	needle = " " + needle

	p := int64(1e9 + 7)
	var tHash int64
	for i := 1; i <= n; i++ {
		tHash = (tHash*131 + (int64)(needle[i]-'a'+1)) % p
	}

	sHash := make([]int64, m+1)
	p131 := make([]int64, m+1)
	p131[0] = 1
	for i := 1; i <= m; i++ {
		sHash[i] = (sHash[i-1]*131 + (int64)(haystack[i]-'a'+1)) % p
		p131[i] = p131[i-1] * 131 % p
	}

	// 滑动窗结尾
	for i := n; i <= m; i++ {
		if calcHash(sHash, p131, p, i-n+1, i) == tHash {
			return i - n
		}
	}

	return -1
}

func calcHash(H, p131 []int64, p int64, l, r int) int64 {
	return ((H[r]-H[l-1]*p131[r-l+1])%p + p) % p
}
