package reverseString2

func reverseString(s []byte) {
	nE := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[nE] = s[nE], s[i]
		nE--
	}
}
