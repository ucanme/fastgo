package strutil

import (
	"bytes"
	"unicode/utf8"
)

func RemoveAllChineseLetters(s string) string {
	return RemoveChineseLetters(s, -1)
}

func RemoveChineseLetters(s string, num int) string {
	if !utf8.ValidString(s) || len(s) == 0 {
		return s
	}

	buf := bytes.Buffer{}
	times := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if (r == utf8.RuneError || size == 3) && (num == -1 || times < num) { // Chinese characters are encoded into 3 bytes in utf-8
			times++
		} else {
			buf.WriteRune(r)
		}
		s = s[size:]
	}

	return buf.String()
}
