package _10_regexp

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/k0kubun/pp"
)

func TestCode(t *testing.T) {
	e := func(s string, p string, expect bool) {
		pp.Printf("s string %v, p string %v, expect bool %v\n", s, p, expect)
		assert.Equal(t, expect, isMatch(s, p))
	}
	e("aa", "a", false)
	e("aa", "aa", true)
	e("aaa", "aa", false)
	e("aaa", "a*a", true)
	e("aaa", "a*aaa", true)
	e("aaa", "a*aaaa", false)
	e("aa", "a*", true)
	e("aa", ".*", true)
	e("ab", ".*", true)
	e("aab", "c*a*b", true)
	e("aab", ".*d*p*b", true)
	e("aab", "d*p*b", false)
	e("aab", "a*d*p*b", true)
	e("aab", "aa.*b", true)
	e("aaa", "ab*a*c*a", true)
}

func BenchmarkCode(t *testing.B) {
	for i := 0; i < t.N; i++ {
		isMatch("aab", ".*d*p*b")
	}
}
