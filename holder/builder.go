package holder

func BuildEmptyField(width, height int) (field Field) {
	items := make([][] Item, height)
	for i := 0; i < height; i++ {
		items[i] = make([] Item, width)
		for j := 0; j < width; j++ {
			items[i][j] = Item{StateUnknown}
		}
	}

	field = Field{items}

	return
}

func BuildEmptyNonogram(width, height int) (nonogram Nonogram) {
	field := BuildEmptyField(width, height)
	leftBox := Box{make([][]int, height)}
	topBox := Box{make([][]int, width)}

	nonogram = Nonogram{field, leftBox, topBox}
	return
}
