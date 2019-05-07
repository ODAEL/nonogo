package holder

type Box struct {
	Numbers [][]int
}

func (box *Box) GetNumbers() [][]int {
	return box.Numbers
}

func (box *Box) SetNumbers(numbers [][]int) {
	box.Numbers = numbers
}

func (box *Box) GetNumbersLine(index int) []int {
	return box.Numbers[index]
}
