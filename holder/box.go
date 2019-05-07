package holder

type Box struct {
	Numbers [][]int
}

func (box *Box) GetNumbersLine(index int) []int {
	return box.Numbers[index]
}
