package solver

import "nonogram-solver/holder"

func findFirstByState(line []*holder.Item, state int) int {
	for i := 0; i < len(line); i++ {
		if line[i].GetState() == state {
			return i
		}
	}

	return -1
}

func findFirstUnknown(line []*holder.Item) int {
	return findFirstByState(line, holder.StateUnknown)
}

func findFirstBlack(line []*holder.Item) int {
	return findFirstByState(line, holder.StateBlack)
}

func findFirstWhite(line []*holder.Item) int {
	return findFirstByState(line, holder.StateWhite)
}

func trimLeftByState(line []*holder.Item, state int) []*holder.Item {
	var i int
	for i = 0; (line[i].GetState() == state) && (i != len(line)-1); i++ {
	}
	return line[i:]
}

func trimRightByState(line []*holder.Item, state int) []*holder.Item {
	var i int
	for i = len(line) - 1; (line[i].GetState() == state) && (i != 0); i-- {
	}
	return line[:i+1]
}

func trimByState(line []*holder.Item, state int) []*holder.Item {
	line = trimLeftByState(line, state)
	line = trimRightByState(line, state)

	return line
}

func trimWhite(line []*holder.Item) []*holder.Item {
	return trimByState(line, holder.StateWhite)
}
