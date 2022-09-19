package longestCommonPrefix

func longestCommonPrefix(strs []string) string {

	var result string
	tmpMap := make(map[uint8]int)
	for i := 0; i < len(strs[0]); i++ {
		for _, str_one := range strs {
			if i >= len(str_one) {
				break
			}
			tmpMap[str_one[i]]++
		}

		if tmpMap[strs[0][i]] != len(strs) {
			break
		}

		result += string(strs[0][i])
		tmpMap = make(map[uint8]int)
	}

	return result
}

func longestCommonPrefix2(strs []string) string {

	var result string
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return result
			}
		}
		result += string(strs[0][i])
	}

	return result
}
