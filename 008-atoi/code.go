package _08_atoi

import "math"

func myAtoi(s string) int {
	converting := false
	result := 0
	base := 10
	minus := false

	for _, r := range s {
		if result >= math.MaxInt32 {
			goto FINAL
		}
		if !converting {
			switch {
			case '0' <= r && r <= '9':
				converting = true
			case r == '-':
				minus = true
				fallthrough
			case r == '+':
				converting = true
				continue
			case r == '0':
				// version only decimal
				return 0
				/* version with 8, 16
				if len(s) <= k + 1 {
					return 0
				}
				if s[k + 1] == 'x' {
					base = 16
				} else {
					base = 8
				}
				converting = true
				continue
				//*/
			case r == ' ':
				continue
			default:
				return 0
			}
		}
		if base == 10 {
			switch  {
			case '0' <= r && r <= '9':
				result *= base
				result += int(r - '0')
				continue
			default:
				goto FINAL
			}
		}
		/* version with 8, 16
		if base == 8 {
			switch  {
			case '0' <= r && r <= '7':
				result *= base
				result += int(r - '0')
				continue
			default:
				goto FINAL
			}
		}
		if base == 16 {
			switch  {
			case '0' <= r && r <= '9':
				result *= base
				result += int(r - '0')
				continue
			case 'a' <= r && r <= 'f':
				result *= base
				result += int(r - 'a' - 10)
				continue
			case 'A' <= r && r <= 'F':
				result *= base
				result += int(r - 'A' - 10)
				continue
			default:
				goto FINAL
			}
		}
		//*/
	}
	FINAL:
	if minus {
		result *= -1
	}
	if math.MaxInt32 < result {
		result = math.MaxInt32
	} else if math.MinInt32 > result {
		result = math.MinInt32
	}
	return result
}
