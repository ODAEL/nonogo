package holder

type Field struct {
	Items [][]Item
}

func (field *Field) GetItems() [][]Item {
	return field.Items
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
