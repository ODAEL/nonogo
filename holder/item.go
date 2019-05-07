package holder

const StateUnknown int = 0
const StateBlack int = 1
const StateWhite int = -1

type Item struct {
	State int
}

func (item *Item) IsUnknown() bool {
	return item.State == StateUnknown
}

func (item *Item) IsBlack() bool {
	return item.State == StateBlack
}

func (item *Item) IsWhite() bool {
	return item.State == StateWhite
}

func (item *Item) PaintUnknown() {
	item.State = StateUnknown
}

func (item *Item) PaintBlack() {
	item.State = StateBlack
}

func (item *Item) PaintWhite() {
	item.State = StateWhite
}
