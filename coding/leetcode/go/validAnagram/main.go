package validAnagram

func isAnagram(s string, t string) bool {

	ALen := len(s)
	A := make(map[int32]int)
	for _, c := range s {
		A[c]++
	}

	for _, c := range t {
		value, ok := A[c]
		if value <= 0 {
			return false
		}
		if ok {
			A[c]--
			ALen--
		}
	}

	if ALen != 0 {
		return false
	}

	return true
}
