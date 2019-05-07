package holder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateTestNumbers() [][]int {
	return [][]int{
		{1, 2, 3},
		{5, 6, 7},
	}
}

func TestBox_GetNumbersLine(t *testing.T) {
	box := Box{generateTestNumbers()}

	assert.Equal(t, generateTestNumbers()[0], box.GetNumbersLine(0))
}
