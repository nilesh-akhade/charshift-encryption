package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(encrypt("hello world", []int{2, 1, 3}))
}

func encrypt(text string, password []int) string {
	var newRune int
	var sb strings.Builder
	for i, rune := range text {
		shift := password[i%len(password)]
		newRune = int(rune) + shift
		if rune >= 97 && rune <= 122 {
			if newRune > 122 {
				newRune = newRune - (122 - 97 + 1)
			}
		} else if rune >= 65 && rune <= 90 {
			if newRune > 90 {
				newRune = newRune - (90 - 65 + 1)
			}
		} else if rune >= 48 && rune <= 57 {
			if newRune > 57 {
				newRune = newRune - (57 - 48 + 1)
			}
		} else {
			sb.WriteRune(rune)
			continue
		}
		fmt.Printf("%d: %d(%c)+%d=%d(%c)\n", i, rune, rune, shift, newRune, newRune)
		sb.WriteString(fmt.Sprintf("%c", newRune))
	}
	return sb.String()
}
