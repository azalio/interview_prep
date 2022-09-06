package implementStrStr

func strStr(haystack string, needle string) int {

	needleLen := len(needle)

	for i := 0; i <= len(haystack)-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
