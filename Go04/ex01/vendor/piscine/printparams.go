package piscine

import (
	"ft"
	"os"
)

func PrintParams() {
	args := os.Args

	for i, arg := range args {
		if i == 0 {
			continue
		}
		for _, r := range arg {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
