package holder

// Nonogram

type Nonogram struct {
	field   Field
	leftBox Box
	topBox  Box
}

func (nonogram *Nonogram) GetField() Field {
	return nonogram.field
}

func (nonogram *Nonogram) GetLeftBox() Box {
	return nonogram.leftBox
}

func (nonogram *Nonogram) GetTopBox() Box {
	return nonogram.topBox
}

func (nonogram *Nonogram) GetHeight() int {
	return nonogram.field.GetHeight()
}

func (nonogram *Nonogram) GetWidth() int {
	return nonogram.field.GetWidth()
}

// Field

type Field struct {
	Items [][] Item
}

func (field *Field) GetHeight() int {
	return len(field.Items)
}

func (field *Field) GetWidth() int {
	if len(field.Items) == 0 {
		return 0
	} else {
		return len(field.Items[0])
	}
}

// Box

type Box struct {
	numbers [][] int
}

func (box *Box) GetNumbers() [][] int {
	return box.numbers
}

func (box *Box) GetNumbersLine(index int) [] int {
	return box.numbers[index]
}

// Item

const StateUnknown int = -1
const StateBlack int = 1
const StateWhite int = 0

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

func (nonogram *Nonogram) GetHorizontalLine(index int) ([] *Item, []int) {
	line := make([] *Item, len(nonogram.field.Items[index]))
	for i := 0; i < len(nonogram.field.Items[index]); i++ {
		line[i] = &nonogram.field.Items[index][i]
	}

	numbers := nonogram.leftBox.GetNumbersLine(index)
	return line, numbers
}

func (nonogram *Nonogram) GetVerticalLine(index int) ([] *Item, []int) {
	line := make([] *Item, len(nonogram.field.Items))
	for i := 0; i < len(nonogram.field.Items); i++ {
		line[i] = &nonogram.field.Items[i][index]
	}

	numbers := nonogram.topBox.GetNumbersLine(index)
	return line, numbers
}