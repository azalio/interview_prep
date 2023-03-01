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

func isValid2(s string) bool {

	pMap := make(map[string]string)
	pMap["]"] = "["
	pMap["}"] = "{"
	pMap[")"] = "("

	var stack []string

	for _, c := range s {
		pResult, inMap := pMap[string(c)]
		if inMap {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == pResult {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(c))
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
