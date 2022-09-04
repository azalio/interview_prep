package stringToInteger_atoi_

import (
	"math"
)

const AsciiToDigit = 48

func myAtoi(s string) int {
	var dResult int
	negative := false
	digitStart := -1
	gotSign := false
	for i := 0; i < len(s); i++ {
		b := s[i]
		if digitStart < 0 {
			if b == ' ' {
				if gotSign {
					return 0
				}
				continue
			}

			if b == '-' {
				if gotSign {
					return 0
				}
				negative = true
				gotSign = true
				continue
			}

			if b == '+' {
				if gotSign {
					return 0
				}
				//negative = false
				gotSign = true
				continue
			}

			if b < '0' || b > '9' {
				return 0
			}

		}

		if b >= '0' && b <= '9' {
			digitStart = i
			d := int(b) - AsciiToDigit

			if negative != true {
				if dResult*10+d > math.MaxInt32 {
					return math.MaxInt32
				}
			} else {
				if (dResult*10+d)*-1 < math.MinInt32 {
					return math.MinInt32
				}
			}

			dResult = dResult*10 + d
			continue
		}

		if digitStart >= 0 {
			break
		}
	}

	if digitStart < 0 {
		return 0
	}

	if negative {
		dResult *= -1
	}

	return dResult
}
