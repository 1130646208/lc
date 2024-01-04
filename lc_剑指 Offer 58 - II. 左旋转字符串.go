package main

import (
	"bytes"
)

func reverseLeftWords(s string, n int) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[(i+n)%len(s)])
	}
	return buf.String()
}
