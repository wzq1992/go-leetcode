package homework

import (
	"reflect"
	"sort"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	testCases := []struct {
		name     string
		in       []string
		expected []string
	}{
		{"empty string slice", []string{}, nil},
		{"empty string", []string{""}, nil},
		{"one section string", []string{"9001"}, nil},
		{"test case 3", []string{"9001 discuss.leetcode.com"}, []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}},
		{"test case 4", []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"1 intel.mail.com", "5 org", "5 wiki.org", "50 yahoo.com", "900 google.mail.com", "901 mail.com", "951 com"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := subdomainVisits(tc.in)
			sort.Strings(got)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("subdomainVisits(%#v) = %#v, want = %#v", tc.in, got, tc.expected)
			}
		})
	}
}
