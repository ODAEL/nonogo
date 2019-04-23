package main

func findFirstByState(line []*Item, state int) int {
	for i := 0; i < len(line); i++ {
		if line[i].GetState() == state {
			return i
		}
	}

	return -1
}

func findFirstUnknown(line []*Item) int {
	return findFirstByState(line, StateUnknown)
}

func findFirstBlack(line []*Item) int {
	return findFirstByState(line, StateBlack)
}

func findFirstWhite(line []*Item) int {
	return findFirstByState(line, StateWhite)
}

func trimLeftByState(line []*Item, state int) []*Item {
	var i int
	for i = 0; (line[i].GetState() == state) && (i != len(line)-1); i++ {
	}
	return line[i:]
}

func trimRightByState(line []*Item, state int) []*Item {
	var i int
	for i = len(line) - 1; (line[i].GetState() == state) && (i != 0); i-- {
	}
	return line[:i+1]
}

func trimByState(line []*Item, state int) []*Item {
	line = trimLeftByState(line, state)
	line = trimRightByState(line, state)

	return line
}

func trimWhite(line []*Item) []*Item {
	return trimByState(line, StateWhite)
}
