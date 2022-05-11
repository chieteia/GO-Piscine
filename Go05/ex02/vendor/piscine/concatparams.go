package piscine

//func strlen(s string) int {
//	var n int
//	for range s {
//		n++
//	}
//	return n
//}

func slicelen(s []string) int {
	var n int
	for range s {
		n++
	}
	return n
}

func ConcatParams(args []string) string {
	var s string
	n := slicelen(args)
	for i, arg := range args {
		s += arg
		if i != n-1 {
			s += "\n"
		}
	}
	return s
}
