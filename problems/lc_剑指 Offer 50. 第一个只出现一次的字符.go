package main

func firstUniqChar(s string) byte {
	hashMap := map[rune]bool{}

	for _, ch := range s {
		if _, ok := hashMap[ch]; !ok {
			hashMap[ch] = true
		} else if ok {
			hashMap[ch] = false
		}
	}

	for _, ch := range s {
		if hashMap[ch] {
			return byte(ch)
		}
	}

	return ' '
}
