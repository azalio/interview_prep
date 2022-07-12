package isomorficStrings

import "fmt"

func isIsomorphic(s string, t string) bool {

	tmpH := make(map[uint8]uint8)

	//"paper",
	//"title",

	for i, v := range t {
		fmt.Printf("s[i]: %d, t[i]: %d\n", s[i], t[i])
		val, ok := tmpH[uint8(v)]

		if !ok {
			fmt.Printf("val: %d\n", val)
			tmpH[uint8(v)] = s[i]
			continue
		}

		if val != s[i] {
			return false
		}
	}

	tmpH2 := make(map[uint8]uint8)

	for i, v := range s {
		fmt.Printf("s[i]: %d, t[i]: %d\n", s[i], t[i])
		val, ok := tmpH2[uint8(v)]

		if !ok {
			fmt.Printf("val: %d\n", val)
			tmpH2[uint8(v)] = t[i]
			continue
		}

		if val != t[i] {
			return false
		}
	}

	//tmpH := make(map[uint8]uint8)
	//tmp2H := make(map[uint8]uint8)
	//
	//  "badc",
	//	"coba",

	//for i, v := range s {
	//	fmt.Printf("s[i]: %d, t[i]: %d\n", s[i], t[i])
	//	val, ok := tmpH[uint8(v)]
	//	val2, ok2 := tmp2H[uint8(v)]
	//
	//	if !ok && !ok2 {
	//		tmpH[uint8(v)] = t[i]
	//		tmp2H[t[i]] = uint8(v)
	//
	//	} else if val != t[i] || val2 != s[i] {
	//		return false
	//	}
	//	fmt.Printf("%v\n", tmpH)
	//	fmt.Printf("%v\n", tmp2H)
	//
	//}
	return true
}
