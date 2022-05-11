package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	arg := os.Args[0]
	// for i, arg := range args {
	// 	if i == 0 {
	for _, r := range arg {
		ft.PrintRune(r)
	}
	// 	}
	// }
	//name := []rune(args[0])
	//for c := range name {
	//	ft.PrintRune(c)
	//}
	ft.PrintRune('\n')
}
