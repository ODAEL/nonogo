package holder

type Box struct {
	numbers [][]int
}

func (box *Box) GetNumbers() [][]int {
	return box.numbers
}

func (box *Box) GetNumbersLine(index int) []int {
	return box.numbers[index]
}
