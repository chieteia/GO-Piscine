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

func RevParams() {
	args := os.Args
	n := slicelen(args)
	for i := n - 1; i > 0; i-- {
		for _, r := range args[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
