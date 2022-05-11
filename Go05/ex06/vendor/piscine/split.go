package piscine

import "fmt"

func strlen(s string) int {
	var n int
	for range s {
		n++
	}
	return n
}

func Split(s, sep string) []string {
	sbytes := append([]byte(s), sep...)
	fmt.Printf("[%v]", string(sbytes))
	fmt.Println([]byte(s))
	//slen, seplen := strlen(s), strlen(sep)
	//split := []string{}

	//prev, cur := 0, 0

	//for cur < slen {

	//	if cur+seplen < slen && s[prev:] == sep {

	//	} else {
	//		en += 1
	//	}

	//	if st+seplen < slen && s[st:en] == sep {
	//		if st != 0 {
	//			split = append(split, s[st:en])
	//			st += seplen
	//		}
	//	} else {
	//		en += 1
	//	}
	//}
	//return split
	return []string{}
}
