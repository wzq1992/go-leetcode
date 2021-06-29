package lessons

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		strSlice := strings.Split(str, "")
		sort.Strings(strSlice)

		sortStr := strings.Join(strSlice, "")
		groups[sortStr] = append(groups[sortStr], str)
	}

	var ans [][]string
	for _, group := range groups {
		ans = append(ans, group)
	}

	return ans
}
