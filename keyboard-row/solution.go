package p500

import "unicode"

func FindWords(words []string) []string {
	keyboard := []string{
		"qwertyuiop", "asdfghjkl", "zxcvbnm",
	}
	keymap := make(map[rune]int)
	for row, item := range keyboard {
		for _, c := range []rune(item) {
			keymap[unicode.ToLower(c)] = row
		}
	}

	res := make([]string, 0)
	for _, word := range words {
		cs := []rune(word)
		if len(cs) == 0 {
			continue
		}
		row := keymap[unicode.ToLower(cs[0])]
		for _, c := range cs {
			if keymap[unicode.ToLower(c)] != row {
				row = -1
				break
			}
		}
		if row != -1 {
			res = append(res, word)
		}
	}
	return res
}
