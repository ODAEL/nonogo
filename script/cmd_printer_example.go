package script

import (
	"nonogram-solver/holder"
	"nonogram-solver/printer"
)

func EmptyNonogramPrintExample() {
	myNonogram := holder.BuildEmptyNonogram(5, 5)
	printer.SimpleCmdPrint(myNonogram)
}