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

	//fmt.Printf("strM: %v\n", strM)

	result := [2]int{-1, -1}

	for _, v := range strM {
		//fmt.Printf("v: %v\n", v)
		if v[1] == 0 {
			if v[0] < result[0] || result[0] == -1 {
				result = v
			}
		}
	}
	return result[0]
}
