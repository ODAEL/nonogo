package holder

type Nonogram struct {
	field        Field
	leftGroup  Group
	rightGroup Group
}

type Field struct {
	items [][] Item
}

type Group struct {
	numbers [][] int
}

const StateUnknown int = -1
const StateBlack int = 1
const StateWhite int = 0

type Item struct {
	state int
}
