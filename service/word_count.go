package service

import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	s = deleteSymbols([]byte(s))
	for _, v := range strings.Split(s, " ") {
		if len(v) != 0 {
			m[v]++
		}
	}
	return m
}

func deleteSymbols(sl []byte) string {
	var s string
	for i, v := range sl {
		if v == 32 || (v > 64 && v < 91) || (v > 96 && v < 123) {
			s += string(v)
		}
		if i != 0 && i != len(sl)-1 && v == '-' && sl[i-1] != ' ' && sl[i+1] != ' ' {
			s += string(v)
		}
	}
	return s
}
