package piscine

import "ft"

func PrintNumber(nb int) {
	if nb < 0 {
		ft.PrintRune('-')
		PrintNumber((nb / 10) * -1)
		ft.PrintRune(rune((nb % 10) * -1 + '0'))
	}
	if nb <= 9 {
		ft.PrintRune(rune(nb + '0'))
	} else {
		PrintNumber(nb / 10)
		ft.PrintRune(rune((nb % 10) + '0'))
	}
}

func PrintString(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

