package solver

import "nonogram-solver/holder"

func Solve(nonogram *holder.Nonogram) {
	for i := 0; i < nonogram.GetWidth(); i++ {
		tryLine(nonogram.GetVerticalLine(i))
	}

	for i := 0; i < nonogram.GetHeight(); i++ {
		tryLine(nonogram.GetHorizontalLine(i))
	}
}

func tryLine(line []*holder.Item, blocks []int) {
	// We need a virtual block with size 1 in the start of line
	// It needed for an auto-recursive try for a first real block
	var virtualLine []holder.Item
	var virtualBlocks []int

	one := holder.Item{}
	one.PaintBlack()
	two := holder.Item{}
	two.PaintWhite()

	virtualLine = append(virtualLine, one, two)

	for i := 0; i < len(line); i++ {
		virtualLine = append(virtualLine, *line[i])
	}

	virtualBlocks = append(virtualBlocks, 1)

	for i := 0; i < len(blocks); i++ {
		virtualBlocks = append(virtualBlocks, blocks[i])
	}

	var virtualCanBlack, virtualCanWhite []bool

	for i := 0; i < len(virtualLine); i++ {
		virtualCanBlack = append(virtualCanBlack, false)
		virtualCanWhite = append(virtualCanWhite, false)
	}

	tryBlock(virtualLine, virtualBlocks, 0, 0, &virtualCanBlack, &virtualCanWhite)

	canBlack := virtualCanBlack[2:]
	canWhite := virtualCanWhite[2:]

	for i := 0; i < len(canBlack); i++ {
		switch true {
		case canWhite[i] && !canBlack[i]:
			line[i].PaintWhite()
			break;
		case canBlack[i] && !canWhite[i]:
			line[i].PaintBlack()
			break;
		case !canBlack[i] && !canWhite[i]:
			// TODO Incorrect nonogram
			break;
		}
	}
}

func tryBlock(
	line []holder.Item,
	blocks []int,
	itemIndex, blockIndex int,
	canBlack, canWhite *[]bool) bool {

	for i := itemIndex; i < blocks[blockIndex]; i++ {
		if line[i].IsWhite() {
			return false
		}
	}

	if blockIndex < len(blocks)-1 {
		result := false

		for nextItemIndex := itemIndex + blocks[blockIndex] + 1; nextItemIndex < len(line)-blocks[blockIndex+1]+2; nextItemIndex++ {
			if line[nextItemIndex-1].IsBlack() {
				break
			}
			if tryBlock(line, blocks, nextItemIndex, blockIndex+1, canBlack, canWhite) {
				result = true
				for i := itemIndex; i < itemIndex+blocks[blockIndex]; i++ {
					(*canBlack)[i] = true
				}
				for i := itemIndex + blocks[blockIndex]; i < nextItemIndex; i++ {
					(*canWhite)[i] = true
				}
			}
		}

		return result

	} else {
		for i := itemIndex + blocks[blockIndex]; i < len(line); i++ {
			if line[i].IsBlack() {
				return false
			}
		}

		for i := itemIndex; i < itemIndex+blocks[blockIndex]; i++ {
			(*canBlack)[i] = true
		}
		for i := itemIndex + blocks[blockIndex]; i < len(line); i++ {
			(*canWhite)[i] = true
		}

		return true
	}
}
