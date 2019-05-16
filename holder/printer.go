package holder

import (
	"fmt"
)

type NonogramPrinter interface {
	Print()
}

type CmdNonogramPrinter struct {
	Nonogram Nonogram
}

func (printer *CmdNonogramPrinter) PrintAll() {
	fmt.Println("Top box")
	simpleCmdBoxPrint(printer.Nonogram.TopBox)

	fmt.Println("Left box")
	simpleCmdBoxPrint(printer.Nonogram.LeftBox)

	fmt.Println("Field")
	simpleCmdFieldPrint(printer.Nonogram.Field)
}

func (printer *CmdNonogramPrinter) PrintTopBox() {
	simpleCmdBoxPrint(printer.Nonogram.TopBox)
}

func (printer *CmdNonogramPrinter) PrintLeftBox() {
	simpleCmdBoxPrint(printer.Nonogram.LeftBox)
}

func (printer *CmdNonogramPrinter) PrintField() {
	simpleCmdFieldPrint(printer.Nonogram.Field)
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
