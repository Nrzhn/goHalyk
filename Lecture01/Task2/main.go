package main

func main() {

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func longestCommonPrefix(strs []string) string {
	s := ""
	if len(strs) == 0 {
		return s
	}

	minL := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minL = min(minL, len(strs[i]))
	}
	for i := 0; i < minL; i++ {
		curr := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != curr {
				return s
			}
		}
		s = s + string(curr)
	}
	return s
}
