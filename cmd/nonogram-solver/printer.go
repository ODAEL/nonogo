package main

import "fmt"

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

	fmt.Println("field")
	simpleCmdFieldPrint(printer.nonogram.GetField())
}

func simpleCmdBoxPrint(box Box) {
	fmt.Println("┏")

	for i := 0; i < len(box.GetNumbers()); i++ {
		fmt.Print("┃")

		for j := 0; j < len(box.GetNumbersLine(i)); j++ {
			fmt.Print(box.GetNumbersLine(i)[j], " ")
		}

		fmt.Println()
	}

	fmt.Println("┗")
}

func simpleCmdFieldPrint(field Field) {
	fmt.Print("┏")
	for i := 0; i < 2*len(field.items[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┓")

	for i := 0; i < len(field.items); i++ {
		fmt.Print("┃")
		for j := 0; j < len(field.items[i]); j++ {
			cmdItemPrint(field.items[i][j])
		}
		fmt.Println("┃")
	}

	fmt.Print("┗")
	for i := 0; i < 2*len(field.items[0]); i++ {
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
		c = "❨❩"
	}

	fmt.Print(c)
}
