package example

import (
	"nonogram-solver/holder"
	"nonogram-solver/printer"
	"nonogram-solver/solver"
)

func EmptyNonogramPrint() {
	myNonogram := holder.BuildEmptyNonogram(5, 5)
	printer.SimpleCmdPrint(myNonogram)
}

func SimpleFulfilledNonogram() {
	boxesNumbers := make([][] int, 2)
	boxesNumbers[0] = make([] int, 1)
	boxesNumbers[1] = make([] int, 1)

	boxesNumbers[0][0] = 2
	boxesNumbers[1][0] = 2

	myNonogram := holder.BuildNonogram(2, 2, boxesNumbers, boxesNumbers)
	solver.Solve(&myNonogram)
	printer.SimpleCmdPrint(myNonogram)
}