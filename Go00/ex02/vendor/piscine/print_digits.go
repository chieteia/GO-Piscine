package piscine

import "ft"

func PrintDigits() {
	for d := '0'; d <= '9'; d++ {
		ft.PrintRune(d)
	}
	ft.PrintRune('\n')
}

