package romanToInteger

var romanian = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int

	for i := 0; i < len(s); {
		if i < len(s)-1 {
			switch s[i] {
			case 'I':
				if s[i+1] == 'V' {
					result += 4
					i += 2
					continue
				} else if s[i+1] == 'X' {
					result += 9
					i += 2
					continue
				}
			case 'X':
				if s[i+1] == 'L' {
					result += 40
					i += 2
					continue
				} else if s[i+1] == 'C' {
					result += 90
					i += 2
					continue
				}
			case 'C':
				if s[i+1] == 'D' {
					result += 400
					i += 2
					continue
				} else if s[i+1] == 'M' {
					result += 900
					i += 2
					continue
				}
			}
		}
		result += romanian[rune(s[i])]
		i++
	}
	return result
}
