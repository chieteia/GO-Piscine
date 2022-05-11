package piscine

import (
	"ft"
	"os"
)

func slicelen(s []string) int {
	var n int
	for range s {
		n++
	}
	return n
}

func SortParams() {
	args := os.Args
	n := slicelen(args)

	if n < 2 {
		return
	}

	args = args[1:]

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-2-i; j++ {
			if args[j+1] < args[j] {
				args[j+1], args[j] = args[j], args[j+1]
			}
		}
	}
	for _, arg := range args {
		for _, r := range arg {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
