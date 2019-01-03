package printer

import (
	"fmt"
	"nonogram-solver/holder"
)

func SimpleCmdPrint(nonogram holder.Nonogram) {

	fmt.Println("Top box")
	simpleCmdBoxPrint(nonogram.TopBox)

	fmt.Println("Left box")
	simpleCmdBoxPrint(nonogram.LeftBox)

	fmt.Println("Field")
	simpleCmdFieldPrint(nonogram.Field)
}

func simpleCmdBoxPrint(box holder.Box) {
	fmt.Println("┏")

	for i := 0; i < len(box.Numbers); i++ {
		fmt.Print("┃")

		for j := 0; j < len(box.Numbers[i]); j++ {
			fmt.Print(box.Numbers[i][j], ' ')
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

	switch item.State {
	case holder.StateBlack:
		c = "██"
	case holder.StateUnknown:
		c = "  "
	case holder.StateWhite:
		c = "❨❩"
	}

	fmt.Print(c)
}
