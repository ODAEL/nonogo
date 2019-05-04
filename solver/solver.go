package solver

import "nonogram-solver/holder"

func Solve(nonogram *holder.Nonogram) {
	for i := 0; i < nonogram.GetWidth(); i++ {
		rulePaintEmpty(nonogram.GetVerticalLine(i))
		ruleSimpleBoxes(nonogram.GetVerticalLine(i))
	}

	for i := 0; i < nonogram.GetHeight(); i++ {
		rulePaintEmpty(nonogram.GetHorizontalLine(i))
		ruleSimpleBoxes(nonogram.GetHorizontalLine(i))
	}
}
