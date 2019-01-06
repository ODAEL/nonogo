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

	topNumbers := [][]int{
		[]int{3, 3},
		[]int{},
	}

	leftNumbers := [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
		[]int{},
		[]int{},
		[]int{1},
		[]int{1},
		[]int{1},
	}

	myNonogram := holder.BuildNonogram(2, 8, leftNumbers, topNumbers)
	solver.Solve(&myNonogram)
	printer.SimpleCmdPrint(myNonogram)
}