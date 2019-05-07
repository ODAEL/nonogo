package holder

type Nonogram struct {
	Field   Field
	LeftBox Box
	TopBox  Box
}

func (nonogram *Nonogram) GetHeight() int {
	return nonogram.Field.GetHeight()
}

func (nonogram *Nonogram) GetWidth() int {
	return nonogram.Field.GetWidth()
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
	line := make([]*Item, len(nonogram.Field.Items[index]))
	for i := 0; i < len(nonogram.Field.Items[index]); i++ {
		line[i] = &nonogram.Field.Items[index][i]
	}

	numbers := nonogram.LeftBox.GetNumbersLine(index)
	return line, numbers
}

func (nonogram *Nonogram) GetVerticalLine(index int) ([]*Item, []int) {
	line := make([]*Item, len(nonogram.Field.Items))
	for i := 0; i < len(nonogram.Field.Items); i++ {
		line[i] = &nonogram.Field.Items[i][index]
	}

	numbers := nonogram.TopBox.GetNumbersLine(index)
	return line, numbers
}
