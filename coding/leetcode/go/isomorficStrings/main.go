package isomorficStrings

func isIsomorphic(s string, t string) bool {

	tmpH := make(map[uint8]uint8)
	tmpH2 := make(map[uint8]uint8)

	for i, v := range t {
		val, ok := tmpH[uint8(v)]

		if !ok {
			tmpH[uint8(v)] = s[i]
			continue
		}

		if val != s[i] {
			return false
		}
	}

	for i, v := range s {
		val, ok := tmpH2[uint8(v)]

		if !ok {
			tmpH2[uint8(v)] = t[i]
			continue
		}

		if val != t[i] {
			return false
		}
	}
	return true
}

func isIsomorphic2(s string, t string) bool {

	tmpH := make(map[uint8]uint8)
	tmpH2 := make(map[uint8]uint8)

	for i := range s {

		val, ok := tmpH[s[i]]
		val2, ok2 := tmpH2[t[i]]

		if ok {
			if val != t[i] {
				return false
			}
		}

		if ok2 {
			if val2 != s[i] {
				return false
			}
		}

		tmpH[s[i]] = t[i]
		tmpH2[t[i]] = s[i]

	}

	return true
}
