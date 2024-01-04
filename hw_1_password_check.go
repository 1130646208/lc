package main

import (
	"fmt"
)

func main() {
	var passwords []string
	var i int
	for {
		var p string
		num, _ := fmt.Scanf("%s", &p)
		if num == 0 {
			break
		}
		passwords = append(passwords, p)
		i++
	}

	for _, p := range passwords {
		if checkPassword(p) {
			fmt.Println("OK")
			return
		}
		fmt.Println("NG")
	}
}

func checkPassword(password string) bool {
	var u, l, d, e, length int

	for i, r := range password {
		length = i + 1

		if r >= 'A' && r <= 'Z' {
			u = 1
		} else if r >= 'a' && r <= 'z' {
			l = 1
		} else if r >= '0' && r <= '9' {
			d = 1
		} else {
			e = 1
		}
	}
	if length < 8 || u+l+d+e < 3 {
		return false
	}

	for i := 0; i < length; i++ {
		for j := i + 2; j < length-2; j++ {
			if password[i] == password[j] && password[i+1] == password[j+1] && password[i+2] == password[j+2] {
				return false
			}
		}
	}

	return true
}
