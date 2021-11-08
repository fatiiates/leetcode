package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	lcp := strs[0]

	for i := 0; i < len(strs); i++ {
		l := 0
		if len(lcp) < len(strs[i]) {
			l = len(lcp)
		} else {
			l = len(strs[i])
		}
		lcp = lcp[:l]

		for j := 0; j < l; j++ {
			if lcp[j] != strs[i][j] {
				lcp = lcp[:j]
				break
			}
		}
		if lcp == "" {
			return ""
		}
	}
	return lcp
}

func main() {

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

}
