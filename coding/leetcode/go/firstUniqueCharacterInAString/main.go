package firstUniqueCharacterInAString

func firstUniqChar(s string) int {

	strM := make(map[int32][2]int)

	for i, c := range s {
		a, ok := strM[c]

		if ok {
			a[1]++
		} else {
			a[0] = i
		}

		strM[c] = a

	}

	result := [2]int{-1, -1}

	for _, v := range strM {
		if v[1] == 0 {
			if v[0] < result[0] || result[0] == -1 {
				result = v
			}
		}
	}
	return result[0]
}

func firstUniqChar2(s string) int {

	strS := make([]int, len(s), len(s))
	strM := make(map[int32]int)

	for i, c := range s {

		a, ok := strM[c]
		if ok {
			strS[a]++
			strS[i] = -1
		} else {
			strM[c] = i
		}
	}

	for i, v := range strS {
		if v == 0 {
			return i
		}
	}

	return -1
}

func firstUniqChar3(s string) int {

	strM := make(map[int32][2]int)

	for i, c := range s {
		a, ok := strM[c]

		if ok {
			a[1]++
		} else {
			a[0] = i
		}

		strM[c] = a

	}

	for i, c := range s {
		v := strM[c]
		if v[1] == 0 {
			return i
		}
	}
	return -1
}
