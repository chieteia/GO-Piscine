package piscine

import "ft"

func PrintIntSlice(buf []int, sep string) {
	for i, n := range buf {
		if i != 0 {
			PrintString(sep)
		}
		if n < 10 {
			ft.PrintRune('0')
		}
		PrintNumber(n)
	}
}

func PrintCombinationNest(from int, until int, n int, buf []int, sep string, flag bool) {
	if n != 0 {
		for i := from; i < until; i++ {
			PrintCombinationNest(i + 1, until, n - 1, append(buf, i), sep, flag && i == from)
		}
	} else {
		if !flag {
			PrintString(", ")
		}
		PrintIntSlice(buf, sep)
	}
}

func PrintCombination(m int, n int, sep string) {
	PrintCombinationNest(0, m, n, make([]int, 0), sep, true);
	ft.PrintRune('\n')
}

