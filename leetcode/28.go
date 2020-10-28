func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+len(needle)-1 >= len(haystack) {
				break
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}