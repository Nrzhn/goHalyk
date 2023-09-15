package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	s := ""
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) > len(s) {
			s = strs[i]
		}
	}
	return s
}
