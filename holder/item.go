package holder

const StateUnknown int = 0
const StateBlack int = 1
const StateWhite int = -1

type Item struct {
	state int
}

func (item *Item) GetState() int {
	return item.state
}

func (item *Item) IsUnknown() bool {
	return item.state == StateUnknown
}

func (item *Item) IsBlack() bool {
	return item.state == StateBlack
}

func (item *Item) IsWhite() bool {
	return item.state == StateWhite
}

func (item *Item) PaintUnknown() {
	item.state = StateUnknown
}

func (item *Item) PaintBlack() {
	item.state = StateBlack
}

func (item *Item) PaintWhite() {
	item.state = StateWhite
}
