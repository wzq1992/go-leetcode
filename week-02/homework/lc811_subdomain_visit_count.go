package homework

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	counters := map[string]int{}

	for _, s := range cpdomains {
		if s == "" {
			continue
		}
		ss := strings.Split(s, " ")
		if len(ss) != 2 {
			continue
		}
		timesStr, fullDomain := ss[0], ss[1]
		times, _ := strconv.Atoi(timesStr)

		counters[fullDomain] += times

		subDomains := strings.Split(fullDomain, ".")
		for i := 1; i < len(subDomains); i++ {
			domain := strings.Join(subDomains[i:], ".")
			counters[domain] += times
		}
	}

	var ans []string
	for domain, times := range counters {
		ans = append(ans, fmt.Sprintf("%d %s", times, domain))
	}

	return ans
}
