package main

import (
	"bytes"
)

func reverseWords(s string) string {
	buffer := bytes.NewBuffer(nil)

	i, j := len(s)-1, len(s)-1
	for j >= 0 {
		if s[j] == ' ' {
			j--
			i = j
		} else {
			for i >= -1 {
				if i == -1 || s[i] == ' ' {
					buffer.WriteString(s[i+1:j+1] + " ")
					j = i
					break
				}
				i--
			}
		}
	}
	res := buffer.String()
	if len(res) > 0 {
		return res[:len(res)-1]
	}
	return ""
}

func main() {
	reverseWords("hello world")
}
