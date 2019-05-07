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
	simpleCmdBoxPrint(printer.nonogram.GetTopBox())

	fmt.Println("Left box")
	simpleCmdBoxPrint(printer.nonogram.GetLeftBox())

	fmt.Println("Field")
	simpleCmdFieldPrint(printer.nonogram.GetField())
}

func simpleCmdBoxPrint(box Box) {
	fmt.Println("┏")

	for i := 0; i < len(box.GetNumbers()); i++ {
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
	for i := 0; i < 2*len(field.GetItems()[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┓")

	for i := 0; i < len(field.GetItems()); i++ {
		fmt.Print("┃")
		for j := 0; j < len(field.GetItems()[i]); j++ {
			cmdItemPrint(field.GetItems()[i][j])
		}
		fmt.Println("┃")
	}

	fmt.Print("┗")
	for i := 0; i < 2*len(field.GetItems()[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┛")
}

func cmdItemPrint(item Item) {
	var c string

	switch item.GetState() {
	case StateBlack:
		c = "██"
	case StateUnknown:
		c = "  "
	case StateWhite:
		c = "><"
	}

	fmt.Print(c)
}
