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

func BuildEmptyNonogram(width, height int) Nonogram {
	leftBoxNumbers := make([][]int, height)
	topBoxNumbers := make([][]int, width)

	return BuildNonogram(width, height, leftBoxNumbers, topBoxNumbers)
}

func BuildNonogram(width, height int, leftBoxNumbers, topBoxNumbrs [][] int) Nonogram {
	field := BuildEmptyField(width, height)
	leftBox := Box{leftBoxNumbers}
	topBox := Box{topBoxNumbrs}

	return Nonogram{field, leftBox, topBox}
}
