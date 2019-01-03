package holder

type Nonogram struct {
	Field   Field
	LeftBox Box
	TopBox  Box
}

type Field struct {
	Items [][] Item
}

type Box struct {
	Numbers [][] int
}

const StateUnknown int = -1
const StateBlack int = 1
const StateWhite int = 0

type Item struct {
	State int
}
