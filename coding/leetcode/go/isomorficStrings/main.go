package isomorficStrings

import "fmt"

func isIsomorphic(s string, t string) bool {

	tmpH := make(map[uint8]uint8)
	tmp2H := make(map[uint8]uint8)

	for i, v := range s {
		fmt.Printf("s[i]: %d, t[i]: %d\n", s[i], t[i])
		val, ok := tmpH[uint8(v)]
		val2, _ := tmp2H[t[i]]

		if !ok {
			tmpH[uint8(v)] = t[i]
			tmp2H[t[i]] = uint8(v)

		} else if val != t[i] || val2 != s[i] {
			return false
		}
		fmt.Printf("%v\n", tmpH)

	}
	return true
}
