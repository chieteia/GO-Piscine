package piscine

func strlen(s []byte) int {
	var n int
	for range s {
		n++
	}
	return n
}

func isWhiteSpace(c rune) bool {
	if c == '\f' || c == '\n' || c == '\r' ||
		c == '\t' || c == '\v' || c == ' ' {
		return true
	} else {
		return false
	}
}

func SplitWhiteSpaces(s string) []string {
	var words []byte
	split := []string{}
	for _, c := range s {
		if isWhiteSpace(c) {
			if strlen(words) > 0 {
				split = append(split, string(words))
				words = []byte{}
			}
		} else {
			words = append(words, byte(c))
		}
	}

	if strlen(words) > 0 {
		split = append(split, string(words))
	}

	return split
}
