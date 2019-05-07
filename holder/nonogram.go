package holder

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

func BuildEmptyNonogram(width, height int) Nonogram {
	leftBoxNumbers := make([][]int, height)
	topBoxNumbers := make([][]int, width)

	return BuildNonogram(width, height, leftBoxNumbers, topBoxNumbers)
}

func BuildNonogram(width, height int, leftBoxNumbers, topBoxNumbers [][]int) Nonogram {
	field := CreateEmptyField(width, height)
	leftBox := Box{leftBoxNumbers}
	topBox := Box{topBoxNumbers}

	return Nonogram{field, leftBox, topBox}
}

func (nonogram *Nonogram) GetHorizontalLine(index int) ([]*Item, []int) {
	line := make([]*Item, len(nonogram.field.items[index]))
	for i := 0; i < len(nonogram.field.items[index]); i++ {
		line[i] = &nonogram.field.items[index][i]
	}

	numbers := nonogram.leftBox.GetNumbersLine(index)
	return line, numbers
}

func (nonogram *Nonogram) GetVerticalLine(index int) ([]*Item, []int) {
	line := make([]*Item, len(nonogram.field.items))
	for i := 0; i < len(nonogram.field.items); i++ {
		line[i] = &nonogram.field.items[i][index]
	}

	numbers := nonogram.topBox.GetNumbersLine(index)
	return line, numbers
}
