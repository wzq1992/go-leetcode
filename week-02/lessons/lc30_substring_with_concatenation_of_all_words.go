package lessons

var wordsMap map[string]int

func findSubstring(s string, words []string) []int {
	wordsMap = countWords(words)
	tot := len(words) * len(words[0])
	var ans []int
	for start := 0; start+tot <= len(s); start++ {
		if isSame(string(s[start:start+tot]), words) {
			ans = append(ans, start)
		}
	}

	return ans
}

// 判断一个字符串t 是否由 words 拼成
// 把 t 分解为若干个单词，然后和 words 比较是否相同（顺序无关）
func isSame(t string, words []string) bool {
	m := len(words[0])
	tMap := map[string]int{}

	for i := 0; i < len(t); i += m {
		tMap[string(t[i:i+m])]++
	}

	return equals(tMap, wordsMap)
}

func equals(a map[string]int, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}

	return true
}

func countWords(words []string) map[string]int {
	ans := map[string]int{}
	for _, word := range words {
		ans[word]++
	}

	return ans
}
