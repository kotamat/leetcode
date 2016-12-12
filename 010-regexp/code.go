package _10_regexp

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		return len(p) >= 2 && p[1] == '*' && isMatch(s, p[2:])
	}
	if len(p) == 1 || p[1] != '*' {
		return (p[0] == '.' || s[0] == p[0]) && isMatch(s[1:], p[1:])
	}
	return isMatch(s, p[2:]) || // p[0] の個数が0個の場合
		(p[0] == '.' || s[0] == p[0]) && isMatch(s[1:], p) // p[0]の個数が1個以上のとき

}
