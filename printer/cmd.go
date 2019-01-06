package printer

import (
	"fmt"
	"nonogram-solver/holder"
)

func SimpleCmdPrint(nonogram holder.Nonogram) {

	fmt.Println("Top box")
	simpleCmdBoxPrint(nonogram.GetTopBox())

	fmt.Println("Left box")
	simpleCmdBoxPrint(nonogram.GetLeftBox())

	fmt.Println("field")
	simpleCmdFieldPrint(nonogram.GetField())
}

func simpleCmdBoxPrint(box holder.Box) {
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

func simpleCmdFieldPrint(field holder.Field) {
	fmt.Print("┏")
	for i := 0; i < 2 * len(field.Items[0]); i++ {
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
	for i := 0; i < 2 * len(field.Items[0]); i++ {
		fmt.Print("━")
	}
	fmt.Println("┛")
}

func cmdItemPrint(item holder.Item) {
	var c string

	switch item.GetState() {
	case holder.StateBlack:
		c = "██"
	case holder.StateUnknown:
		c = "  "
	case holder.StateWhite:
		c = "❨❩"
	}

	fmt.Print(c)
}
