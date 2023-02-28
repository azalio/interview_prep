package validParentheses

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	if string(s[0]) == ")" || string(s[0]) == "}" || string(s[0]) == "]" {
		return false
	}

	var ptr int
	var pipe []uint8
	pipe = append(pipe, s[0])

	for _, c := range s[1:] {
		if len(pipe) == 0 {
			pipe = append(pipe, uint8(c))
			ptr = 0
			continue
		}
		if string(pipe[ptr]) == "(" {
			if string(c) == ")" {
				pipe = pipe[:ptr]
				ptr--
			} else if string(c) == "[" || string(c) == "{" || string(c) == "(" {
				pipe = append(pipe, uint8(c))
				ptr++
			} else if string(c) == "]" || string(c) == "}" {
				return false
			}
			continue
		}

		if string(pipe[ptr]) == "{" {
			if string(c) == "}" {
				pipe = pipe[:ptr]
				ptr--
			} else if string(c) == "[" || string(c) == "{" || string(c) == "(" {
				pipe = append(pipe, uint8(c))
				ptr++
			} else if string(c) == "]" || string(c) == ")" {
				return false
			}
			continue
		}

		if string(pipe[ptr]) == "[" {
			if string(c) == "]" {
				pipe = pipe[:ptr]
				ptr--
			} else if string(c) == "[" || string(c) == "{" || string(c) == "(" {
				pipe = append(pipe, uint8(c))
				ptr++
			} else if string(c) == "}" || string(c) == ")" {
				return false
			}
			continue
		}
	}

	if len(pipe) != 0 {
		return false
	}

	return true
}
