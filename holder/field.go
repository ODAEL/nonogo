package holder

type Field struct {
	Items [][]Item
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

func (field *Field) GetFieldAsSlice() [][]int {
	var slice [][]int

	height := field.GetHeight()
	width := field.GetWidth()

	for i := 0; i < height; i++ {
		var row []int
		for j := 0; j < width; j++ {
			row = append(row, field.Items[i][j].State)
		}
		slice = append(slice, row)
	}

	return slice
}

func CreateEmptyField(width, height int) Field {
	items := make([][]Item, height)
	for i := 0; i < height; i++ {
		items[i] = make([]Item, width)
		for j := 0; j < width; j++ {
			items[i][j] = Item{StateUnknown}
		}
	}

	return Field{items}
}

func CreateFieldBySlice(fieldSlice [][]int) Field {
	var fieldItems [][]Item

	for i := 0; i < len(fieldSlice); i++ {
		var lineItems []Item

		for j := 0; j < len(fieldSlice[i]); j++ {
			lineItems = append(lineItems, Item{fieldSlice[i][j]})
		}

		fieldItems = append(fieldItems, lineItems)
	}

	return Field{fieldItems}
}
