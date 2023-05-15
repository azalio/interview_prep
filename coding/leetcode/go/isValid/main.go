package isValid

import "fmt"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

func openOrClosed(s string) bool {
	openBrackets := []string{"(", "[", "{"}

	for _, char := range openBrackets {
		if s == char {
			return true
		}
	}
	return false
}

func isValid(s string) bool {

	brackets := make(map[string]string)
	brackets[")"] = "("
	brackets["]"] = "["
	brackets["}"] = "{"

	var stk []string

	for _, char := range s {
		strChar := fmt.Sprintf("%c", char)
		if openOrClosed(strChar) {
			stk = append(stk, strChar)
		} else if len(stk) > 0 && stk[len(stk)-1] == brackets[strChar] {
			stk = stk[:len(stk)-1]
		} else {
			return false
		}
	}

	if len(stk) > 0 {
		return false
	}
	return true
}
