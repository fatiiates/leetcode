package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := [][]string{}
	sort.Strings(strs)
	for _, v := range strs {
		str := strings.Split(v, "")
		sort.Strings(str)
		sorted_str := strings.Join(str, "")

		if _, ok := m[sorted_str]; ok {
			res[m[sorted_str]] = append(res[m[sorted_str]], v)
		} else {
			res = append(res, []string{v})
			m[sorted_str] = len(res) - 1
		}
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
