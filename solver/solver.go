package solver

import "nonogram-solver/holder"

func Solve(nonogram *holder.Nonogram) {
	for i := 0; i < nonogram.GetWidth(); i++ {
		ruleSimpleBoxes(nonogram.GetVerticalLine(i))
	}

	for i := 0; i < nonogram.GetHeight(); i++ {
		ruleSimpleBoxes(nonogram.GetHorizontalLine(i))
	}
}