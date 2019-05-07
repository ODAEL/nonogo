package holder

import (
	"fmt"
)

type NonogramPrinter interface {
	Print()
}

type CmdNonogramPrinter struct {
	nonogram Nonogram
}

func (printer *CmdNonogramPrinter) Print() {
	fmt.Println("Top box")
	simpleCmdBoxPrint(printer.nonogram.TopBox)

	fmt.Println("Left box")
	simpleCmdBoxPrint(printer.nonogram.LeftBox)

	fmt.Println("Field")
	simpleCmdFieldPrint(printer.nonogram.Field)
}

func simpleCmdBoxPrint(box Box) {
	fmt.Println("┏")

	for i := 0; i < len(box.Numbers); i++ {
		fmt.Print("┃")

		for j := 0; j < len(box.GetNumbersLine(i)); j++ {
			fmt.Print(box.GetNumbersLine(i)[j])
			if j != len(box.GetNumbersLine(i)) - 1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	fmt.Println("┗")
}

func simpleCmdFieldPrint(field Field) {
	fmt.Print("┏")
	for i := 0; i < 2*len(field.Items[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┓")

	for i := 0; i < len(field.Items); i++ {
		fmt.Print("┃")
		for j := 0; j < len(field.Items[i]); j++ {
			cmdItemPrint(field.Items[i][j])
		}
		fmt.Println("┃")
	}

	fmt.Print("┗")
	for i := 0; i < 2*len(field.Items[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┛")
}

func cmdItemPrint(item Item) {
	var c string

	switch item.State {
	case StateBlack:
		c = "██"
	case StateUnknown:
		c = "  "
	case StateWhite:
		c = "><"
	}

	fmt.Print(c)
}
