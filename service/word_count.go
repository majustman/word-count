package service

import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	s = replaceWhiteSpaces(s)
	for _, v := range strings.Split(s, " ") {
		if len(v) != 0 {
			m[v]++
		}
	}
	return m
}

func replaceWhiteSpaces(s string) string {
	b := []byte(s)
	for i, v := range b {
		if v == '\n' || v == '\t' {
			b[i] = ' '
		}
	}
	return string(b)
}
