package main

import (
	"bytes"
)

func replaceSpace(s string) string {
	var buf bytes.Buffer
	for _, ch := range s {
		if ch == ' ' {
			buf.Write([]byte("%20"))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
