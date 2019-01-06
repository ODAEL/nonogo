package solver

import "nonogram-solver/holder"

func ruleSimpleBoxes(line [] *holder.Item, numbers [] int) {
	line = trimWhite(line)

	if len(numbers) == 0 {
		return
	}

	if len(numbers) > 1 {
		firstBarLength := numbers[0]
		lastBarLength := numbers[len(numbers) - 1]

		ruleSimpleBoxes(line[firstBarLength + 1:], numbers[1:])
		ruleSimpleBoxes(line[:len(line) - lastBarLength - 1], numbers[:len(numbers) - 1])
	}

	// len(numbers) == 1
	barLength := numbers[0]
	lineLength := len(line)

	// If lineLength - barLength > barLength no iterations would made
	for i := lineLength - barLength; i < barLength; i++ {
		line[i].PaintBlack()
	}
}

func rulePaintEmpty(line [] *holder.Item, numbers [] int) {
	if len(numbers) != 0 {
		return
	}

	for i := 0; i < len(line); i++ {
		line[i].PaintWhite()
	}
}