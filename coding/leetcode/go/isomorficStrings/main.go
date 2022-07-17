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

func forString(s, t string, ch, quit chan bool) {

	tmpH := make(map[uint8]uint8)

	for i, v := range t {
		select {
		case <-quit:
			ch <- false
			return
		default:
		}

		val, ok := tmpH[uint8(v)]

		if !ok {
			tmpH[uint8(v)] = s[i]
			continue
		}

		if val != s[i] {
			ch <- false
			quit <- true
			return
		}
	}

	ch <- true
	return
}

func isIsomorphic3(s string, t string) bool {

	quit := make(chan bool)
	ch := make(chan bool)
	ch2 := make(chan bool)

	go forString(t, s, ch, quit)
	go forString(s, t, ch2, quit)

	var s1, s2 bool
	select {
	case s1 = <-ch:
		if s1 != true {
			return false
		} else {
			s2 = <-ch2
		}
	case s2 = <-ch2:
		if s2 == false {
			return false
		} else {
			s1 = <-ch
		}
	}

	if s1 == true && s2 == true {
		return true
	}

	return false
}
